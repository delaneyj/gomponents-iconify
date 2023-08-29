package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdRewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M249.6 402V110L32 256l217.6 146zm12.8-146L480 402V110L262.4 256z" fill="currentColor"/>`),
		g.Group(children),
	)
}