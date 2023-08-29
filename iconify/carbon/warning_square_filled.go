package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningSquareFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.002 4H5.998A1.998 1.998 0 0 0 4 5.998v20.004A1.998 1.998 0 0 0 5.998 28h20.004A1.998 1.998 0 0 0 28 26.002V5.998A1.998 1.998 0 0 0 26.002 4ZM14.875 8h2.25v10h-2.25ZM16 24a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 24Z"/><path fill="none" d="M14.875 8h2.25v10h-2.25ZM16 24a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 24Z"/>`),
		g.Group(children),
	)
}