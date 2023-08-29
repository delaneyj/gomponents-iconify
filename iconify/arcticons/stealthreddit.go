package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stealthreddit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.855 31.923l9.39-18.221l-18.169 9.442"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.82 41.047L42.5 5.5L6.953 20.179l-1.421 1.449l-.032 1.978l10.576-.462l-1.107 1.579l-.043 4.7l-4.516 2.336l-2.69 2.774l7.405-1.661l-1.66 7.407l2.775-2.69l2.334-4.517l4.701-.041l1.579-1.108l-.461 10.577l1.978-.032l1.45-1.42Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.927 29.423l.2 3.45l3.448.2"/>`),
		g.Group(children),
	)
}