package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FuelLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 13.087c0-2.096 0-3.145.553-3.94c.554-.795 1.524-1.14 3.465-1.833l8-2.855c3.636-1.297 5.454-1.946 6.718-1.031C22 4.342 22 6.305 22 10.233v5.65c0 2.884 0 4.326-.879 5.221C20.243 22 18.828 22 16 22H8c-2.828 0-4.243 0-5.121-.896C2 20.21 2 18.767 2 15.884v-2.797Z"/><path d="M9 14c0-1.414 0-2.121.44-2.56C9.878 11 10.585 11 12 11c1.414 0 2.121 0 2.56.44c.44.439.44 1.146.44 2.56c0 1.414 0 2.121-.44 2.56c-.439.44-1.146.44-2.56.44c-1.414 0-2.121 0-2.56-.44C9 16.122 9 15.415 9 14Z"/><path stroke-linecap="round" d="m15 11l1-1m-7 1l-1-1m7 7l1 1m-7-1l-1 1M5 7.06c0-1.305 0-1.957.338-2.407c.087-.116.189-.22.302-.308C6.08 4 6.72 4 8 4h.818c.507 0 .761 0 .97.057a1.653 1.653 0 0 1 1.156 1.18" opacity=".5"/></g>`),
		g.Group(children),
	)
}