package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CwFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCw4x30"><path fill-opacity=".7" d="M0 0h682.7v512H0z"/></clipPath><path id="flagCw4x31" d="m0-1l.2.7H1L.3 0l.2.7L0 .4l-.6.4l.2-.7l-.5-.4h.7z"/></defs><g clip-path="url(#flagCw4x30)" transform="scale(.94)"><path fill="#002b7f" d="M0 0h768v512H0z"/><path fill="#f9e814" d="M0 320h768v64H0z"/><use width="13500" height="9000" x="2" y="2" fill="#fff" href="#flagCw4x31" transform="scale(42.67)"/><use width="13500" height="9000" x="3" y="3" fill="#fff" href="#flagCw4x31" transform="scale(56.9)"/></g>`),
		g.Group(children),
	)
}