package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M4 7v6h8V7H4zm2 5H5v-1h1v1zm0-2H5V9h1v1zm3 2H7v-1h2v1zm0-2H7V9h2v1zm2 2h-1v-1h1v1zm0-2h-1V9h1v1z"/>`),
		g.Group(children),
	)
}