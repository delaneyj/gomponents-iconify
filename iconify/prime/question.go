package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Question(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.07 4.93A5.75 5.75 0 0 0 6.25 9a.75.75 0 0 0 1.5 0A4.26 4.26 0 1 1 12 13.25a.76.76 0 0 0-.75.75v2a.75.75 0 0 0 1.5 0v-1.3a5.76 5.76 0 0 0 3.32-9.77Z"/><circle cx="12" cy="19.5" r="1.25" fill="currentColor"/>`),
		g.Group(children),
	)
}