package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11h6m-6 8v-1.25c0-2.071-1.919-3.75-4.286-3.75H8.286C5.919 14 4 15.679 4 17.75V19m9-11a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/>`),
		g.Group(children),
	)
}