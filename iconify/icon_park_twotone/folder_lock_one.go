package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderLockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFolderLockOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M43 23v-9a2 2 0 0 0-2-2H24l-5-6H7a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h15"/><path fill="#555" d="M29 34h14v8H29z"/><path d="M39 34v-3a3 3 0 1 0-6 0v3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFolderLockOne0)"/>`),
		g.Group(children),
	)
}