package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Types(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 17a5 5 0 1 1-10.001-.001A5 5 0 0 1 22 17zM6.5 6.5h3.8L7 1L1 11h5.5V6.5zm9.5 4.085V8H8v8h2.585A6.505 6.505 0 0 1 16 10.585z"/>`),
		g.Group(children),
	)
}