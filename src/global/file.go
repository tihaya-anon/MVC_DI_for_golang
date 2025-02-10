package global

var PATH = Path{
	RESOURCE: Resource{
		TEMPLATE: Template{
			CONTROLLER: Controller{
				DIR:     []string{"..", "resource", "template", "controller"},
				BUILDER: []string{"controller_builder.txt"},
				CORE:    []string{"controller_core.txt"},
				ROUTER:  []string{"controller_router.txt"},
			},
			MAPPER: Mapper{
				DIR:       []string{"..", "resource", "template", "mapper"},
				INTERFACE: []string{"mapper.txt"},
				IMPL:      []string{"mapper_impl.txt"},
			},
			SERVICE: Service{
				DIR:       []string{"..", "resource", "template", "service"},
				INTERFACE: []string{"service_interface.txt"},
				BUILDER:   []string{"service_builder.txt"},
				IMPL:      []string{"service_impl.txt"},
			},
		},
	},
	CONTROLLER: Controller{
		BUILDER: []string{"controller", "builder"},
		CORE:    []string{"controller"},
	},
	MAPPER: Mapper{
		INTERFACE: []string{"mapper"},
		IMPL:      []string{"mapper", "impl"},
	},
	SERVICE: Service{
		INTERFACE: []string{"service"},
		BUILDER:   []string{"service", "builder"},
		IMPL:      []string{"service", "impl"},
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
}

type Mapper = struct {
	DIR       []string
	INTERFACE []string
	IMPL      []string
}

type Service = struct {
	DIR       []string
	BUILDER   []string
	INTERFACE []string
	IMPL      []string
}
