package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitTurnDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSplitTurnDownLeft0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M37 22H21a8 8 0 0 0-8 8v14"/><circle cx="37" cy="8.944" r="5" fill="#555" transform="rotate(-90 37 8.944)"/><path stroke-linecap="round" stroke-linejoin="round" d="M37 14v29m5-4l-5 5l-5-5m-14 0l-5 5l-5-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSplitTurnDownLeft0)"/>`),
		g.Group(children),
	)
}