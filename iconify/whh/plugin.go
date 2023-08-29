package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plugin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 768q-35 0-64-17.5T786 704h-18v256q0 27-19 45.5t-45 18.5H448v-18q29-17 46.5-46t17.5-64q0-53-37.5-90.5T384 768t-90.5 37.5T256 896q0 35 17.5 64t46.5 46v18H64q-26 0-45-18.5T0 960V704h18q17 29 46 46.5t64 17.5q53 0 90.5-37.5T256 640t-37.5-90.5T128 512q-35 0-64 17.5T18 576H0V320q0-26 19-45t45-19h256v-18q-29-17-46.5-46T256 128q0-53 37.5-90.5T384 0t90.5 37.5T512 128q0 35-17.5 64T448 238v18h256q26 0 45 19t19 45v256h18q17-29 46-46.5t64-17.5q53 0 90.5 37.5T1024 640t-37.5 90.5T896 768z"/>`),
		g.Group(children),
	)
}