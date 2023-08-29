package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageTrolley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M11.5 5v11m8-11v11m-1-11a3 3 0 1 0-6 0m-8 16.5a2 2 0 1 0 4 0m11 0a2 2 0 1 0 4 0M24 19H4.5v-3.249c0-3.165-.058-6.353-.795-9.432C3.035 3.52 1.877.5 0 .5M23.5 16h-16v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L7.5 5.176V5h16v.176l-.203 1.346a26.749 26.749 0 0 0 0 7.956l.203 1.345V16Z"/>`),
		g.Group(children),
	)
}