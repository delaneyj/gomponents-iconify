package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImplantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20 10a4 4 0 0 1 8 0v28a4 4 0 0 1-8 0V10Zm4-2a2 2 0 0 0-2 2v28a2 2 0 1 0 4 0V10a2 2 0 0 0-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}