package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RubySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.293 4L8.058.236L9.73 4H4.293Zm10-4l-3.632 3.632L9.047 0h5.246ZM.236 8.058L4 9.73V4.293L.236 8.058Zm3.396 2.603L0 9.047v5.246l3.632-3.632ZM5 9.293L9.293 5H5v4.293Zm10 4.438l-3.907-9.117L15 .707v13.024Zm-.952.317l-3.717-8.672l-4.955 4.955l8.672 3.717Zm-9.434-2.955L13.731 15H.707l3.907-3.907Z"/>`),
		g.Group(children),
	)
}