package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListTwoHundredSeventyTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.5 20.75a1.25 1.25 0 1 1 2.499 0a1.25 1.25 0 0 1-2.499 0Zm.5-3.5V2.75a.75.75 0 0 1 1.493-.102l.007.102v14.5a.75.75 0 0 1-1.493.102L18 17.25Zm-7 3.5a1.25 1.25 0 1 1 2.499 0a1.25 1.25 0 0 1-2.499 0Zm.5-3.5V2.75a.75.75 0 0 1 1.493-.102L13 2.75v14.5a.75.75 0 0 1-1.493.102l-.007-.102Zm-7 3.5a1.25 1.25 0 1 1 2.499 0a1.25 1.25 0 0 1-2.499 0Zm.5-3.5V2.75a.75.75 0 0 1 1.493-.102l.007.102v14.5a.75.75 0 0 1-1.493.102L5 17.25Z"/>`),
		g.Group(children),
	)
}