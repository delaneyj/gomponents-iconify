package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoDimensionalCodeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTwoDimensionalCodeOne0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" stroke-linejoin="round" d="M18 6H6v12h12V6Zm0 24H6v12h12V30Zm24 0H30v12h12V30Zm0-24H30v12h12V6Z"/><path stroke-linecap="round" d="M24 6v18m0 6v12m0-18H6m36 0H30"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTwoDimensionalCodeOne0)"/>`),
		g.Group(children),
	)
}