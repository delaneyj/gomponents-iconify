package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jiosaavn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 31.85c4.56-3 14.76 1.08 15.07 10.65c0 0 1.4-23.69-15.07-23.69"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26 5.5c3.93 15.34.75 24.43-3.61 31.31c.42-9.59-7.57-27.13-10.52-31.31"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 27.56c-4.37.2-12.17 4.54-16 14.94c0-12.51 2.23-18.72 16-23.69"/>`),
		g.Group(children),
	)
}