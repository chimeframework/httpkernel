package httpkernel

const (

	//project/src/src dirs
	SRC_DIR                           = "src"
	SRC_BUNDLE_SUFFIX                 = "Bundle"
	SRC_BUNDLE_RESOURCES_DIR          = "Resources"
	SRC_BUNDLE_VIEWS_DIR              = "views"
	SRC_BUNDLE_VIEWS_DEFAULT_DIR      = "Default"
	SRC_BUNDLE_VIEWS_DEFAULT_TEMPLATE = "index"
	SRC_BUNDLE_VIEWS_DEFAULT_LAYOUT   = "layout"

	// project/app dirs
	APP_DIR                            = "app"
	APP_CONFIG_DIR                     = "config"
	APP_CONFIG_FILE                    = "config.yml"
	APP_RESOURCES_DIR                  = "Resources"
	APP_RESOURCES_VIEWS_DIR            = "views"
	APP_RESOURCES_VIEWS_DEFAULT_LAYOUT = "base"

	CONFIG_FRAMEWORK                 = ".framework"
	CONFIG_ROUTER                    = ".router"
	CONFIG_RESOURCE                  = ".resource"
	KERNEL_ROOT_DIR                  = "%kernel.root_dir%"
	KERNEL_APP_DIR                   = "%kernel.app_dir%"
	CONFIG_FRAMEWORK_ROUTER_RESOURCE = CONFIG_FRAMEWORK + CONFIG_ROUTER + CONFIG_RESOURCE
	CONTROLLER_SUFFIX                = "Controller"
	ACTION_SUFFIX                    = "Action"
	SEPARATOR                        = ":"
)
