package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M18.5 8.5h-5.281l-.709-1.24A2.5 2.5 0 0 0 10.34 6H7.5A2.5 2.5 0 0 0 5 8.5V18a2.5 2.5 0 0 0 2.5 2.5h11A2.5 2.5 0 0 0 21 18v-7a2.5 2.5 0 0 0-2.5-2.5Zm-11 11A1.5 1.5 0 0 1 6 18V8.5A1.5 1.5 0 0 1 7.5 7h2.84a1.5 1.5 0 0 1 1.302.756l.852 1.492a.5.5 0 0 0 .435.252H18.5A1.5 1.5 0 0 1 20 11v7a1.5 1.5 0 0 1-1.5 1.5h-11Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}