package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessKingFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2a1 1 0 0 1 .993.883L13 3v2h2a1 1 0 0 1 .117 1.993L15 7h-2v1.758a4.49 4.49 0 0 1 2.033-.734l.24-.018L15.5 8a4.5 4.5 0 0 1 4.5 4.5a4.504 4.504 0 0 1-4.064 4.478l-.217.016L15.5 17h-7a4.5 4.5 0 1 1 2.501-8.241L11 7H9a1 1 0 0 1-.117-1.993L9 5h2V3a1 1 0 0 1 1-1zm6 16H6a1 1 0 0 0-1 1a2 2 0 0 0 2 2h10a2 2 0 0 0 1.987-1.768l.011-.174A1 1 0 0 0 18 18z"/></g>`),
		g.Group(children),
	)
}