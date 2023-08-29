package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserAsking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14H9.286C6.919 14 5 15.679 5 17.75V19M19 7v5a2 2 0 0 1-2 2h-2v5M14 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/>`),
		g.Group(children),
	)
}