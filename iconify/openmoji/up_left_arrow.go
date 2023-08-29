package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="m37.836 14.039l-24.229-.897l.897 24.228l5.558-.205l-.524-14.16L55.43 58.897l3.933-3.932L23.47 19.072l14.16.525z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m37.836 14.039l-24.229-.897l.897 24.228l5.558-.205l-.524-14.16L55.43 58.897l3.933-3.932L23.47 19.072l14.16.525z"/>`),
		g.Group(children),
	)
}