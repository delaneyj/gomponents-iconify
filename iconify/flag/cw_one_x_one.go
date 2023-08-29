package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CwOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCw1x10"><path fill-opacity=".7" d="M0 0h9000v9000H0z"/></clipPath><path id="flagCw1x11" d="m0-1l.2.7H1L.3 0l.2.7L0 .4l-.6.4l.2-.7l-.5-.4h.7z"/></defs><g clip-path="url(#flagCw1x10)" transform="scale(.057)"><path fill="#002b7f" d="M0 0h13500v9000H0z"/><path fill="#f9e814" d="M0 5625h13500v1125H0z"/><use width="13500" height="9000" x="2" y="2" fill="#fff" href="#flagCw1x11" transform="scale(750)"/><use width="13500" height="9000" x="3" y="3" fill="#fff" href="#flagCw1x11" transform="scale(1000)"/></g>`),
		g.Group(children),
	)
}