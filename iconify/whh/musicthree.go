package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musicthree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-265-888q-10-6-29.5-6.5t-37.5 5.5q-9-7-20-7q-13 0-22.5 9.5t-9.5 22.5v513q-55-33-128-33q-80 0-136 37.5t-56 90.5t56 90.5t135.5 37.5t136-37.5t56.5-90.5V406q12-22 32-22q33 0 57.5 11t38 27t22 38t11 40t2 38t-1 27t-1.5 11q-32 0-32 32q0 14 8.5 23t23.5 9q13 0 28.5-13t31-37t26-66.5t10.5-94.5q0-37-41-104t-88.5-120.5t-71.5-68.5z"/>`),
		g.Group(children),
	)
}