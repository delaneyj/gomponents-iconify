package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camscanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 5a6 6 0 0 0-6 6v26a6 6 0 0 0 6 6h26a6 6 0 0 0 6-6V11a6 6 0 0 0-6-6ZM5 34.53h38"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.81 27A4.75 4.75 0 0 0 31 28.83h2.51a4.24 4.24 0 0 0 4.24-4.24h0a4.24 4.24 0 0 0-4.24-4.25h-2.8a4.25 4.25 0 0 1-4.24-4.25h0a4.24 4.24 0 0 1 4.24-4.25h2.52a4.73 4.73 0 0 1 4.16 1.86m-15.86 9.43v.07a5.63 5.63 0 0 1-5.63 5.63h0a5.63 5.63 0 0 1-5.63-5.63v-5.73a5.63 5.63 0 0 1 5.63-5.63h0a5.63 5.63 0 0 1 5.63 5.63v.07"/>`),
		g.Group(children),
	)
}