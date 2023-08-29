package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pulse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16c-120 0-135 105-60 195c0-165 135-45 135-135c0-30-45-60-75-60zm146.25 134.532C370.61 152.554 334.75 167.875 301 196c165 0 45 135 135 135c30 0 60-45 60-75c0-75-41.016-108.838-93.75-105.468zM76 181c-30 0-60 45-60 75c0 120 105 135 195 60c-165 0-45-135-135-135zm175.782 15A60 60 0 0 0 196 256a60 60 0 0 0 120 0a60 60 0 0 0-64.218-60zM316 301c0 165-135 45-135 135c0 30 45 60 75 60c120 0 135-105 60-195z"/>`),
		g.Group(children),
	)
}