package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVerticalSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDistributeVerticalSpacing0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path d="M5 8h38M5 40h38"/><path fill="#fff" stroke-linejoin="round" d="M14 20h20v8H14z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDistributeVerticalSpacing0)"/>`),
		g.Group(children),
	)
}