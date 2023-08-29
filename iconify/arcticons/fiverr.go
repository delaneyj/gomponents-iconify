package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fiverr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.709 4.5h-7.474c-5.447 0-10.198 4.294-9.88 12.076H7.99v7.245h5.724V43.5h8.498V23.821h8.855V43.5h8.945V16.576H22.748v-1.879a2.805 2.805 0 0 1 2.848-2.951h5.113Z"/>`),
		g.Group(children),
	)
}