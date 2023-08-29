package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderQuestion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.57 16.3a.64.64 0 0 0-.15-.13l-.17-.09l-.19-.08a1 1 0 0 0-.9.28a1 1 0 0 0-.22.32a1 1 0 0 0-.07.39a1 1 0 0 0 .29.7a1 1 0 0 0 .32.22a1 1 0 0 0 .39.07a1 1 0 0 0 .38-.07a1 1 0 0 0 .32-.22a1 1 0 0 0 .3-.7a1 1 0 0 0-.08-.39a.87.87 0 0 0-.22-.3Zm-.7-7.3a3 3 0 0 0-2.6 1.5a1 1 0 1 0 1.73 1a1 1 0 0 1 1.87.5a1 1 0 0 1-1 1a1 1 0 1 0 0 2a3 3 0 0 0 0-6ZM19 5.5h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3Zm1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}