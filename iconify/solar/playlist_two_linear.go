package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 16.429C16 17.849 14.88 19 13.5 19S11 17.849 11 16.429c0-1.42 1.12-2.572 2.5-2.572S16 15.008 16 16.43Zm0 0V10"/><path stroke-linecap="round" d="m18.675 7.116l-1.755.842c-.341.164-.512.246-.636.373a1.001 1.001 0 0 0-.223.356C16 8.853 16 9.042 16 9.42c0 .877 0 1.315.191 1.578a1 1 0 0 0 .63.396c.319.058.714-.132 1.504-.51l1.755-.843c.341-.164.512-.246.636-.372a1.01 1.01 0 0 0 .223-.356C21 9.147 21 8.957 21 8.578c0-.876 0-1.315-.191-1.578a1 1 0 0 0-.63-.395c-.319-.059-.714.13-1.504.51ZM15 5H3m10 4H3m6 4H3m5 4H3"/></g>`),
		g.Group(children),
	)
}