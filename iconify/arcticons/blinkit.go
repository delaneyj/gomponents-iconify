package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blinkit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="19.157" cy="19.045" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.157 22.345v6.596"/><circle cx="34.84" cy="19.045" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.84 22.345v6.596M15.504 19.17v8.55c0 .733.489 1.221 1.222 1.221h.366m21.026-8.55v7.329c0 .733.489 1.221 1.222 1.221h.366m-2.931-6.596h2.687m-13.485 6.596v-4.03a2.45 2.45 0 0 0-2.443-2.444h0a2.45 2.45 0 0 0-2.442 2.443m0 4.031v-6.596M8.706 24.788a2.45 2.45 0 0 1 2.443-2.443h0a2.45 2.45 0 0 1 2.443 2.443v1.588a2.45 2.45 0 0 1-2.443 2.443h0a2.45 2.45 0 0 1-2.443-2.443m0 2.565V19.17m19.183-.123v9.894m0-2.076l4.397-4.52m-2.931 3.054l3.42 3.42"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}