package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 2.332v26.5h26.5V2.33H2zm24.5 24.5H4V12.5h8.167V4.332H26.5v22.5zM15.63 17.65l5.47 5.468l-1.21 1.206l5.483 1.47l-1.47-5.482l-1.195 1.195l-5.467-5.466l1.21-1.207l-5.482-1.47l1.468 5.48l1.195-1.194z"/>`),
		g.Group(children),
	)
}