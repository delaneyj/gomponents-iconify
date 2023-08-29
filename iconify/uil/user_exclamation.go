package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.55 12.22a4.92 4.92 0 0 0 1.7-3.72a5 5 0 0 0-10 0A4.92 4.92 0 0 0 8 12.22a8 8 0 0 0-4.7 7.28a1 1 0 0 0 2 0a6 6 0 0 1 12 0a1 1 0 0 0 2 0a8 8 0 0 0-4.75-7.28Zm-3.3-.72a3 3 0 1 1 3-3a3 3 0 0 1-3 3Zm8.5-5a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1ZM19 11.79a1.05 1.05 0 0 0-.29.71a1 1 0 0 0 .29.71a1.15 1.15 0 0 0 .33.21a.94.94 0 0 0 .76 0a.9.9 0 0 0 .54-.54a.84.84 0 0 0 .08-.38a1 1 0 0 0-1.71-.71Z"/>`),
		g.Group(children),
	)
}