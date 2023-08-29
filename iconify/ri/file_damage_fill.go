package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDamageFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3 14l4 2.5l3-3.5l3 4l2-2.5l3 .5l-3-3l-2 2.5l-3-5l-3.5 3.75L3 10V2.992C3 2.455 3.447 2 3.998 2H14v6a1 1 0 0 0 1 1h6v11.992A1 1 0 0 1 20.007 22H3.993A.993.993 0 0 1 3 21.008V14Zm18-7h-5V2.003L21 7Z"/>`),
		g.Group(children),
	)
}