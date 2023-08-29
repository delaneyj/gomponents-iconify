package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sberbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.21 7.85L16 20.62L5.21 14.39m29.13-7.83L16 17.13l-8.58-4.88m33.66-2.34L16 24.35L3.85 17.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.39 15.46A19.36 19.36 0 0 1 44.5 22a20.25 20.25 0 0 1-1.11 6.64l-.5 1.36a20.87 20.87 0 0 1-4.38 6.54A20.23 20.23 0 0 1 32 40.87a20.19 20.19 0 0 1-8 1.66a19.93 19.93 0 0 1-8-1.66a21 21 0 0 1-6.49-4.33A19.74 19.74 0 0 1 5.11 30a20.71 20.71 0 0 1-1.61-8v-1.36L16 27.73l26.29-15.09Z"/>`),
		g.Group(children),
	)
}