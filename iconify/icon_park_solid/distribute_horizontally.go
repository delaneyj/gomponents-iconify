package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDistributeHorizontally0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M30 10v28H18V10z"/><path stroke-linecap="round" d="M40 6v36M8 6v36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDistributeHorizontally0)"/>`),
		g.Group(children),
	)
}