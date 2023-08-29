package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screwdriveralt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M948.5 948.5q-75.5 75.5-183 75.5T583 948L480 845q-1-70-50.5-120T309 672q-16-20-21-42v-35q7-26 32-51l224-224q25-25 51-32h35q21 5 42 21q2 71 52.5 120.5T845 480l103 103q76 75 76 182.5t-75.5 183zm-427-235q-9.5 9.5-9.5 23t10 23.5l222 222q9 10 23 10t23.5-9.5T800 959t-10-23L568 714q-10-10-23.5-10t-23 9.5zM664 618q-10-10-23.5-10t-23 9.5t-9.5 23t10 23.5l254 254q9 10 23 10t23.5-9.5T928 895t-10-23zm96-96q-10-10-23.5-10t-23 9.5t-9.5 23t10 23.5l222 222q9 10 23 10t23.5-9.5T992 767t-10-23zm-408-74L192 288H96L0 96L96 0l192 96v96l160 160z"/>`),
		g.Group(children),
	)
}