package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#9A4E1C" d="M32 8h-6V4a4 4 0 0 0-4-4h-8a4 4 0 0 0-4 4v4H4a4 4 0 0 0-4 4v20a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V12a4 4 0 0 0-4-4zM12 6a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v2H12V6z"/><path fill="#662113" d="M36 20a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4v-8a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v8z"/><path fill="#9A4E1C" d="M36 18a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4v-6a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v6z"/><path fill="#CCD6DD" d="M22 18a2 2 0 0 1-2 2h-4a2 2 0 0 1 0-4h4a2 2 0 0 1 2 2"/>`),
		g.Group(children),
	)
}