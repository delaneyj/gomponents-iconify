package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="m14.504 34.702l-.897 24.229l24.229-.897l-.206-5.558L23.47 53l35.893-35.892l-3.933-3.933l-35.892 35.893l.524-14.16z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m14.504 34.702l-.897 24.229l24.229-.897l-.206-5.558L23.47 53l35.893-35.892l-3.933-3.933l-35.892 35.893l.524-14.16z"/>`),
		g.Group(children),
	)
}