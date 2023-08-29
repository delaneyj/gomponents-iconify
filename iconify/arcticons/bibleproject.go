package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bibleproject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.661 5.5v31.411l8.414-2.89V42.5l10.178-10.177l18.086 4.524V5.5l-18.202 4.981L5.661 5.5z"/>`),
		g.Group(children),
	)
}