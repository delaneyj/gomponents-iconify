package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AldiTalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.875 10.499L14.5 27.501m10.297-17.002l-6.375 17.002m10.297-17.002l-6.375 17.002m7.117-12.752h.871M27.867 19h4.039m-5.632 4.251H33.5"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.065 33.04v5.373h2.687M15.791 33.04h3.56m-1.78 5.373V33.04m11.75 0v5.373m2.888 0l-2.212-2.686l2.212-2.669m-2.212 2.669h-.676m-9.563 2.67l1.745-5.357m1.748 5.373l-1.748-5.373m1.163 3.576h-2.327"/>`),
		g.Group(children),
	)
}