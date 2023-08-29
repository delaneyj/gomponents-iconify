package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileExclamationAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 10h1a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2Zm6 2H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Zm0 4H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Zm4.71 4.29a1.58 1.58 0 0 0-.15-.12a.76.76 0 0 0-.18-.09L19.2 20a1 1 0 0 0-.58.06a.9.9 0 0 0-.54.54a.84.84 0 0 0-.08.4a1 1 0 1 0 1.92-.38a1.15 1.15 0 0 0-.21-.33ZM20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a1 1 0 0 0 0-2H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v2a1 1 0 0 0 2 0V8.94ZM15 8a1 1 0 0 1-1-1V5.41L16.59 8Zm4 7a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}