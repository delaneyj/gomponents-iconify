package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IudOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 6a1 1 0 1 1 0-2h12a1 1 0 0 1 1 1a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H25v30.126A4.002 4.002 0 0 1 24 44a4 4 0 0 1-1-7.874V6H11Zm11 34a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}