package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontalSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDistributeHorizontalSpacing0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path d="M8 43V5m32 38V5"/><path fill="#fff" stroke-linejoin="round" d="M20 14h8v20h-8z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDistributeHorizontalSpacing0)"/>`),
		g.Group(children),
	)
}