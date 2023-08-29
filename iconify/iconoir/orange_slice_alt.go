package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrangeSliceAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.39 10.11L5.61 2.334c-4.295 4.296-4.295 11.26 0 15.556c4.296 4.296 11.26 4.296 15.557 0l-7.778-7.778Zm0 0l.353 8.133m-.354-8.132H5.612m7.779 0l-5.304 5.303"/>`),
		g.Group(children),
	)
}