package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myfritz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40 5.5H7a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.94 36.32v-9.29l4.65 9.3l4.65-9.29v9.29m4.65 0l-2.48-6.16m4.66 0l-3 8.37a1.39 1.39 0 0 1-1.31.93h-.39M5 24h37M23.46 10.79v7.92M10 10.79h3.96M10 14.75h2.57M10 10.79v7.92m6.08 0v-7.92h2.59a2.66 2.66 0 0 1 0 5.32h-2.59m2.59 0l2.59 2.6m4.35-7.92h5.24m-2.62 7.92v-7.92m4.53 0H38l-5.24 7.92H38"/>`),
		g.Group(children),
	)
}