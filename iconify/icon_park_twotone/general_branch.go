package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GeneralBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGeneralBranch0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M40 9H8a2 2 0 0 0-2 2v30a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V11a2 2 0 0 0-2-2Z"/><path stroke-linecap="round" d="M15 5v4m18-4v4M6 17h36M18 30h12m-6-6v12M6 11v12m36-12v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGeneralBranch0)"/>`),
		g.Group(children),
	)
}