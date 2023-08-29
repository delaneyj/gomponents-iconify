package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HostLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityHostLine0" fill="currentColor"><path d="M26.5 2h-17A1.5 1.5 0 0 0 8 3.5V34h20V3.5A1.5 1.5 0 0 0 26.5 2ZM26 32H10V4h16Z"/><path d="M12 6.2h12v1.6H12zm0 4h12v1.6H12zm6 12.58a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4.5a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5Z"/></g>`),
		g.Group(children),
	)
}