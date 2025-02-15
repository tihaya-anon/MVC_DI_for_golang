package global

var PATH = Path{
	RESOURCE: Resource{
		TEMPLATE: Template{
			CONTROLLER: Controller{
				DIR:     []string{"resource", "template", "controller"},
				BUILDER: []string{"controller_builder.txt"},
				CORE:    []string{"controller_core.txt"},
				ROUTER:  []string{"controller_router.txt"},
				TEST:    []string{"controller_test.txt"},
			},
			MAPPER: Mapper{
				DIR:       []string{"resource", "template", "mapper"},
				INTERFACE: []string{"mapper_interface.txt"},
				BUILDER:   []string{"mapper_builder.txt"},
				IMPL:      []string{"mapper_impl.txt"},
			},
			SERVICE: Service{
				DIR:       []string{"resource", "template", "service"},
				INTERFACE: []string{"service_interface.txt"},
				BUILDER:   []string{"service_builder.txt"},
				IMPL:      []string{"service_impl.txt"},
				TEST:      []string{"service_test.txt"},
			},
		},
	},
	CONTROLLER: Controller{
		BUILDER: []string{"controller", "builder"},
		CORE:    []string{"controller"},
		TEST:    []string{"controller", "test"},
	},
	MAPPER: Mapper{
		INTERFACE: []string{"mapper"},
		BUILDER:   []string{"mapper", "builder"},
		IMPL:      []string{"mapper", "impl"},
	},
	SERVICE: Service{
		INTERFACE: []string{"service"},
		BUILDER:   []string{"service", "builder"},
		IMPL:      []string{"service", "impl"},
		TEST:      []string{"service", "test"},
	},
}

type Path = struct {
	RESOURCE   Resource
	CONTROLLER Controller
	MAPPER     Mapper
	SERVICE    Service
}

type Resource = struct {
	TEMPLATE Template
}

type Template = struct {
	CONTROLLER Controller
	MAPPER     Mapper
	SERVICE    Service
}

type Controller = struct {
	DIR     []string
	BUILDER []string
	CORE    []string
	ROUTER  []string
	TEST    []string
}

type Mapper = struct {
	DIR       []string
	INTERFACE []string
	BUILDER   []string
	IMPL      []string
}

type Service = struct {
	DIR       []string
	BUILDER   []string
	INTERFACE []string
	IMPL      []string
	TEST      []string
}
