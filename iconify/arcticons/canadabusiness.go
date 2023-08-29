package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Canadabusiness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 6.247v22.167l-6.479-6.478l-1.588-10.163l3.72 2.096L24 6.247zm0 0v22.167l6.479-6.478l1.588-10.163l-3.72 2.096L24 6.247zM12.948 19.903l10.16 10.037l-9.207 11.813l.762-3.874L4.5 30.066l2.604-1.207l-1.524-6.923h5.843l1.525-2.033zm22.104 0l-21.151 21.85h20.198l-.762-3.874L43.5 30.066l-2.604-1.207l1.524-6.923h-5.843l-1.525-2.033z"/>`),
		g.Group(children),
	)
}