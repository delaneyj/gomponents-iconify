package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Time(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 717C160 717 0 557 0 359S160 0 358 0s359 161 359 359s-161 358-359 358zm0-628C210 89 90 211 90 359s120 268 268 268s270-120 270-268S506 89 358 89zm97 407l-127-93c-9-7-18-22-18-34V166c0-11 10-20 22-20h42c11 0 20 9 20 20v160c0 12 9 27 18 34l93 67c9 7 11 21 4 30l-25 34c-7 9-20 11-29 5z"/>`),
		g.Group(children),
	)
}