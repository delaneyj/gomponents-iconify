package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Irma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.929 18.915v8.654m10.535-.01v-8.644l4.328 8.654l4.327-8.641v8.641m-16.789 0v-8.654h2.833a2.906 2.906 0 0 1 0 5.813H18.33m2.833 0l2.833 2.839m13.294-.024l3.149-8.628m1.535 5.759h-3.543"/><rect width="34.374" height="23.148" x="6.813" y="12.426" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" transform="rotate(-20 24 24)"/>`),
		g.Group(children),
	)
}