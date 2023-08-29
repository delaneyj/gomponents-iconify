package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="m35.134 14.039l24.229-.897l-.897 24.228l-5.558-.205l.524-14.16L17.54 58.897l-3.933-3.932L49.5 19.072l-14.16.525z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m35.134 14.039l24.229-.897l-.897 24.228l-5.558-.205l.524-14.16L17.54 58.897l-3.933-3.932L49.5 19.072l-14.16.525z"/>`),
		g.Group(children),
	)
}