package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageLockers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M8.5 3.5h1M15 5V3.5M13 5V3.5m-5 9v11m8-11v11m-1-11a3 3 0 1 0-6 0M11.599 5A2.999 2.999 0 0 1 6 3.5A3 3 0 0 1 11.599 2h4.813c.347.492.815.885 1.36 1.142L18 3.25v.5l-.229.108A3.486 3.486 0 0 0 16.411 5H11.6ZM20 23.5H4v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L4 12.677V12.5h16v.177l-.203 1.345a26.747 26.747 0 0 0 0 7.956L20 23.323v.177Z"/>`),
		g.Group(children),
	)
}