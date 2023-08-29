package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSearchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFolderSearchOne0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M43 23v-9a2 2 0 0 0-2-2H24l-5-6H7a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h15"/><circle cx="32" cy="32" r="6" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="m37 36l5 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFolderSearchOne0)"/>`),
		g.Group(children),
	)
}