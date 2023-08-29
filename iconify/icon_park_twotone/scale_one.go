package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTScaleOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M18 14H5V9l13-5h12l13 5v5H30"/><path fill="#555" d="M18 4h12v40H18z"/><path stroke-linecap="round" d="M18 12h4m-4 18h5m-5-12h5m-5 6h4m-4 12h4m-4-26v28"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTScaleOne0)"/>`),
		g.Group(children),
	)
}