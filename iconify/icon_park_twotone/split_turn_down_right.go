package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitTurnDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSplitTurnDownRight0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M13 22h16a8 8 0 0 1 8 8v14"/><circle cx="13" cy="8.944" r="5" fill="#555" transform="rotate(-90 13 8.944)"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 14v29m5-4l-5 5l-5-5m34 0l-5 5l-5-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSplitTurnDownRight0)"/>`),
		g.Group(children),
	)
}