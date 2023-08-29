package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.07 6.12A3.48 3.48 0 0 1 12 6a4 4 0 0 1 4 4v1.34a1 1 0 0 0 2 0V10a6 6 0 0 0-5-5.91V3a1 1 0 0 0-2 0v1.1l-.45.08a1 1 0 0 0 .52 1.94Zm10.64 14.17l-18-18a1 1 0 0 0-1.42 1.42l4.12 4.11A6 6 0 0 0 6 10v3.18A3 3 0 0 0 4 16v2a1 1 0 0 0 1 1h3.14a4 4 0 0 0 7.72 0h1.73l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM8 10a3.31 3.31 0 0 1 0-.55L11.59 13H8Zm4 10a2 2 0 0 1-1.72-1h3.44A2 2 0 0 1 12 20Zm-6-3v-1a1 1 0 0 1 1-1h6.59l2 2Z"/>`),
		g.Group(children),
	)
}