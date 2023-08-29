package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func I(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 12a2 2 0 0 1 2-2h16a2 2 0 1 1 0 4h-6v20h6a2 2 0 1 1 0 4H16a2 2 0 1 1 0-4h6V14h-6a2 2 0 0 1-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}