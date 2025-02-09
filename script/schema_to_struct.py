import json
import pathlib


class StructGenerator:
    def __generate_go_struct(self, schema, struct_name):
        if schema.get("type") == "object":
            properties = schema.get("properties", {})
            required_fields = schema.get("required", [])
            struct_def = f"type {struct_name} struct {{\n"
            for prop_name, prop_schema in properties.items():
                field_type = self.__get_go_type(prop_schema)
                field_name = self.__snake_to_pascal(prop_name)
                if prop_name in required_fields:
                    struct_def += f"    {field_name} {field_type}\n"
                else:
                    struct_def += f"    {field_name} *{field_type}\n"
            struct_def += "}\n"
            return struct_def
        return ""
    @staticmethod
    def __snake_to_pascal(s):
        parts = s.split("_")
        return "".join([part.capitalize() for part in parts])
    @staticmethod
    def __get_go_type(schema):
        if "$ref" in schema:
            ref = schema["$ref"].split("/")[-1]
            return ref.lower()
        if schema.get("type") == "string":
            return "string"
        if schema.get("type") == "integer":
            return "int"
        if schema.get("type") == "object":
            return "struct"
        return "interface{}"

    def __generate_nested_structs(self, definitions, name):
        nested_structs = ""
        for def_name, def_schema in definitions.items():
            if def_name != name:
                nested_structs += self.__generate_go_struct(def_schema, def_name.lower()) + "\n"
        return nested_structs

    def generate_file(self, file: str, output: str):
        with open(file, "r") as f:
            schema = json.load(f)
        
        ref = schema.get("$ref", None)
        if not ref:
            print("No $ref found in schema")
            return
        
        name = ref.split("/")[-1]
        definitions = schema.get("definitions", {})

        nested_structs = self.__generate_nested_structs(definitions, name)
        main_struct = self.__generate_go_struct(definitions[name], name)
        output = pathlib.Path(output)
        pkg = output.name
        output = output.joinpath(f"{name.lower()}.go")
        with open(output, "w") as f:
            f.write(f"package {pkg}\n\n")
            f.write(nested_structs)
            f.write(main_struct)
    
    def generate_batch(self, dir:str, output:str):
        files = pathlib.Path(dir).iterdir()
        for idx, file in enumerate(files):
            print(f"Generating {file.name} ({idx+1})")
            self.generate_file(file, output)

if __name__ == "__main__":
    StructGenerator().generate_batch("../schema", "../src/config/model")