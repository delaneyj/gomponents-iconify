package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageCheckIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1 23.5h20.5v-19m0 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-16 4v11m8-11v11m-1-11a3 3 0 1 0-6 0m11 11h-16v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L1.5 8.676V8.5h16v.176l-.203 1.346a26.747 26.747 0 0 0 0 7.956l.203 1.345v.177Z"/>`),
		g.Group(children),
	)
}