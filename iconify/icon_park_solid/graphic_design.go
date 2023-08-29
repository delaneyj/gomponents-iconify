package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicDesign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGraphicDesign0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M19 32c-7.732 0-14-6.268-14-14S11.268 4 19 4s14 6.268 14 14"/><path fill="#fff" d="M44 18H18v26h26V18Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGraphicDesign0)"/>`),
		g.Group(children),
	)
}