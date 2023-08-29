package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListRotateTwoHundredSeventyTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 2.75a.75.75 0 0 0-1.5 0v11.5a.75.75 0 0 0 1.5 0V2.75ZM4.5 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm5.5 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm6.5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM10 2a.75.75 0 0 1 .75.75v11.5a.75.75 0 0 1-1.5 0V2.75A.75.75 0 0 1 10 2Zm6.25.75a.75.75 0 0 0-1.5 0v11.5a.75.75 0 0 0 1.5 0V2.75Z"/>`),
		g.Group(children),
	)
}