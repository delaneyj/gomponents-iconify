package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeHundredSixtyFiveScores(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M20.325 28.727c.785.658 1.546.96 3.45.96h.41a3.674 3.674 0 0 0 3.674-3.674h0a3.674 3.674 0 0 0-3.674-3.674h-3.86v-4.025h7.1"/><circle cx="34.733" cy="25.919" r="3.767"/><path d="M38.167 19.7c-.628-.821-1.584-1.386-3.17-1.386h-.264a3.767 3.767 0 0 0-3.767 3.767v3.838M9.501 28.727c.785.658 1.633.96 3.537.96h1.154a2.842 2.842 0 0 0 2.842-2.843h0A2.842 2.842 0 0 0 14.192 24M9.5 19.264c.787-.656 1.635-.955 3.539-.95l1.153.002a2.842 2.842 0 0 1 2.842 2.843h0A2.842 2.842 0 0 1 14.192 24m-2.896.001h2.896"/></g>`),
		g.Group(children),
	)
}