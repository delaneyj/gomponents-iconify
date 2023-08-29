package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M17 2a3 3 0 0 1 2.995 2.824L20 5a3 3 0 0 1 3 3a6 6 0 0 1-5.775 5.996L17 14h-4l.15.005a2 2 0 0 1 1.844 1.838L15 16v4a2 2 0 0 1-1.85 1.995L13 22h-2a2 2 0 0 1-1.995-1.85L9 20v-4a2 2 0 0 1 1.85-1.995L11 14v-1a1 1 0 0 1 .883-.993L12 12h5a4 4 0 0 0 4-4a1 1 0 0 0-.883-.993L20 7l-.005.176a3 3 0 0 1-2.819 2.819L17 10H7a3 3 0 0 1-2.995-2.824L4 7V5a3 3 0 0 1 2.824-2.995L7 2h10z"/></g>`),
		g.Group(children),
	)
}