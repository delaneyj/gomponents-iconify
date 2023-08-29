package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.92 16.12a.76.76 0 0 0-.09-.18a1.58 1.58 0 0 0-.12-.15l-.15-.12l-.18-.09a.6.6 0 0 0-.19-.06a1 1 0 0 0-.9.27l-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33a1 1 0 0 0 1.09.22a1.46 1.46 0 0 0 .33-.22a1.46 1.46 0 0 0 .22-.33a1 1 0 0 0 .05-.38a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.08-.18ZM12 10.5a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1Zm7-5h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3Zm1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}