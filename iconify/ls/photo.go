package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M522 686V114H0v572h522zM47 551V161h428v390H47zM666 8v571H548V445h70V55H190v38h-47V8h523z"/>`),
		g.Group(children),
	)
}