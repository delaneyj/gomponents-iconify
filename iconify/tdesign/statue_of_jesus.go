package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatueOfJesus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 1h4v4h-4V1ZM2 6h20v2.66l-7 3V17h.78l1.5 6H6.72l1.5-6H9v-5.34l-7-3V6Zm7 3.483V8H5.539L9 9.483ZM11 8v2.586l2 2V8h-2Zm4 0v1.483L18.461 8H15Zm-2 7.414l-2-2V17h2v-1.586ZM14.22 19H9.78l-.5 2h5.44l-.5-2Z"/>`),
		g.Group(children),
	)
}