package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="m57.466 34.702l.897 24.229l-24.229-.897l.206-5.558L48.5 53L12.607 17.108l3.933-3.933l35.892 35.893l-.524-14.16z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m57.466 34.702l.897 24.229l-24.229-.897l.206-5.558L48.5 53L12.607 17.108l3.933-3.933l35.892 35.893l-.524-14.16z"/>`),
		g.Group(children),
	)
}