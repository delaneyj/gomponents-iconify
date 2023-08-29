package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2a4 4 0 0 1 4 4a5.03 5.03 0 0 1-.438 2.001L16 8a1 1 0 0 1 .117 1.993L16 10h-1.263l1.24 5.79a1 1 0 0 1-.747 1.184l-.113.02L15 17H9a1 1 0 0 1-.996-1.093l.018-.117L9.262 10H8a1 1 0 0 1-.117-1.993L8 8h.438a5.154 5.154 0 0 1-.412-1.525l-.02-.259L8 6a4 4 0 0 1 4-4zm6 16H6a1 1 0 0 0-1 1a2 2 0 0 0 2 2h10a2 2 0 0 0 1.987-1.768l.011-.174A1 1 0 0 0 18 18z"/></g>`),
		g.Group(children),
	)
}