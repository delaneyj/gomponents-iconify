package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderProtectionOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFolderProtectionOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M43 23v-9a2 2 0 0 0-2-2H24l-5-6H7a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h15"/><path fill="#555" d="M28 30.8c0-.933 6-2.8 6-2.8s6 1.867 6 2.8c0 7.467-6 11.2-6 11.2s-6-3.733-6-11.2Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFolderProtectionOne0)"/>`),
		g.Group(children),
	)
}