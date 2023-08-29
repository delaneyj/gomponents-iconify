package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceContentArchiveFolderOutboxContentFolderArchiveFileInbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 9.54a1 1 0 0 0-1-1h-3v1h-5v-1h-3a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1Zm-12-4h11m-11-3h11"/>`),
		g.Group(children),
	)
}