package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h8a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2zM4 4v24h8V4zm24 26h-8a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h8a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2zM20 4v24h8V4z"/>`),
		g.Group(children),
	)
}