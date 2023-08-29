package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GripperTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 0q132 0 248 50t204 138t137 203t51 249v768q0 132-50 248t-138 204t-203 137t-249 51q-132 0-248-50t-204-138t-137-203t-51-249V640q0-132 50-248t138-204T775 51t249-51zm0 1024q79 0 149-30t122-83t82-122t31-149q0-79-30-149t-83-122t-122-82t-149-31q-79 0-149 30t-122 83t-82 122t-31 149q0 79 30 149t83 122t122 82t149 31zm0-651q55 0 103 21t85 57t58 85t21 104q0 55-21 103t-57 85t-85 58t-104 21q-55 0-103-21t-85-57t-58-85t-21-104q0-55 21-103t57-85t85-58t104-21z"/>`),
		g.Group(children),
	)
}