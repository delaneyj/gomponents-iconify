package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderWithdrawal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFolderWithdrawal0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M5 8a2 2 0 0 1 2-2h12l5 6h17a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8Z"/><path stroke-linecap="round" d="m21 23l-5 5l5 5"/><path stroke-linecap="round" d="M16 28h16v-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFolderWithdrawal0)"/>`),
		g.Group(children),
	)
}