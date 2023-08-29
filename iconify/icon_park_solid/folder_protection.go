package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderProtection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFolderProtection0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M5 8a2 2 0 0 1 2-2h12l5 6h17a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><path fill="#000" stroke="#000" stroke-linecap="round" d="M18 22.8c0-.933 6-2.8 6-2.8s6 1.867 6 2.8c0 7.467-6 11.2-6 11.2s-6-3.733-6-11.2Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFolderProtection0)"/>`),
		g.Group(children),
	)
}