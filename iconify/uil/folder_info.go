package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13.5a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1Zm7-8h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3Zm1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Zm-7.29-7.71a1 1 0 0 0-1.09-.21a.93.93 0 0 0-.33.21a1.05 1.05 0 0 0-.29.71a1 1 0 0 0 .29.7a1 1 0 0 0 1.42 0a1 1 0 0 0 .29-.7a1 1 0 0 0-.08-.38a.93.93 0 0 0-.21-.33Z"/>`),
		g.Group(children),
	)
}