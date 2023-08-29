package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScalesThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.998 2v1.278l5 1.668l3.633-1.21l.632 1.896l-3.031 1.011l3.095 8.512A5.983 5.983 0 0 1 17.998 17a5.983 5.983 0 0 1-4.328-1.845l3.094-8.512l-3.766-1.256V19h4v2h-10v-2h4V5.387L7.232 6.643l3.095 8.512A5.983 5.983 0 0 1 6 17a5.983 5.983 0 0 1-4.33-1.845l3.095-8.512l-3.03-1.01l.632-1.898L6 4.945l4.999-1.667V2h2Zm5 7.103L16.581 13h2.835l-1.418-3.897Zm-12 0L4.582 13h2.835L5.999 9.103Z"/>`),
		g.Group(children),
	)
}