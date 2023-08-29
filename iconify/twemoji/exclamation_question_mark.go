package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationQuestionMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#BE1931" d="M21 27a3 3 0 0 1-3-3v-4a3 3 0 0 1 3-3c.603-.006 6-1 6-5c0-2-2-4-5-4c-2.441 0-4 2-4 3a3 3 0 1 1-6 0c0-4.878 4.58-9 10-9c8 0 11 5.982 11 11c0 4.145-2.277 7.313-6.413 8.92c-.9.351-1.79.587-2.587.747V24a3 3 0 0 1-3 3z"/><circle cx="21" cy="32" r="3" fill="#BE1931"/><circle cx="6" cy="32" r="3" fill="#BE1931"/><path fill="#BE1931" d="M9 24a3 3 0 1 1-6 0V5a3 3 0 1 1 6 0v19z"/>`),
		g.Group(children),
	)
}