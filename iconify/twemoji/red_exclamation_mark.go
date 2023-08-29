package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="32" r="3" fill="#BE1931"/><path fill="#BE1931" d="M21 24a3 3 0 0 1-6 0V5a3 3 0 1 1 6 0v19z"/>`),
		g.Group(children),
	)
}