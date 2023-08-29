package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketsInfos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.38 25.8v-8H24c1.48 0 2.68 1.2 2.68 2.69s-1.2 2.69-2.68 2.69h-2.62m2.62-.01l2.62 2.63m-9.64-4c1.1 0 2 .9 2 2s-.9 2-2 2h-3.3v-8h3.3c1.1 0 2 .9 2 2s-.9 2-2 2h0Zm0 0h-3.3m18.64 0c1.1 0 2 .9 2 2s-.9 2-2 2h-3.3v-8h3.3c1.1 0 2 .9 2 2s-.9 2-2 2h0Zm0 0h-3.3m-15.34 8.4h20.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/>`),
		g.Group(children),
	)
}