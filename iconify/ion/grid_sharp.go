package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M240 240H32V32h208Zm240 0H272V32h208ZM240 480H32V272h208Zm240 0H272V272h208Z"/>`),
		g.Group(children),
	)
}