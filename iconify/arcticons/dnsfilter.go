package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dnsfilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5s-11.26 2-15.25 2v20a11.16 11.16 0 0 0 .8 4.1a15 15 0 0 0 2 3.61a22 22 0 0 0 2.81 3.07a34.47 34.47 0 0 0 3 2.48a34 34 0 0 0 2.89 1.86c1 .59 1.71 1 2.13 1.19l1 .49a1.44 1.44 0 0 0 1.24 0l1-.49c.42-.2 1.13-.6 2.13-1.19a34 34 0 0 0 2.89-1.86a34.47 34.47 0 0 0 3-2.48a22 22 0 0 0 2.81-3.07a15 15 0 0 0 2-3.61a11.16 11.16 0 0 0 .8-4.1v-20c-3.99.03-15.25-2-15.25-2Z"/><circle cx="24" cy="22.14" r="11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.38 29.25A17 17 0 0 0 24 27.31a17 17 0 0 0-8.38 1.94m0-14.25A17 17 0 0 0 24 17a17 17 0 0 0 8.38-2"/><ellipse cx="24" cy="22.14" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.07" ry="11"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 11.14v22m-11-11h22"/>`),
		g.Group(children),
	)
}