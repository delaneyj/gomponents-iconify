package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SidePanelCloseFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 4H4c-1.1 0-2 .9-2 2v20c0 1.1.9 2 2 2h24c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm0 11H17.8l3.6-3.6L20 10l-6 6l6 6l1.4-1.4l-3.6-3.6H28v9H12V6h16v9z"/>`),
		g.Group(children),
	)
}