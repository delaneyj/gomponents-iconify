package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10m6-6v.172a2 2 0 0 1-.586 1.414l-3.828 3.828a2 2 0 0 1-1.414.586H15m6-6h-4a2 2 0 0 0-2 2v4M7 7h10M7 11h10M7 15h4"/>`),
		g.Group(children),
	)
}