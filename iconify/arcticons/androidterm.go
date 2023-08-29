package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Androidterm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95ZM9 9l8.22 4.74L9 18.48m17.55 0h-9.34M11.42 22a2.57 2.57 0 0 1 1.76.74l2.09 2.08a14.9 14.9 0 0 1 17.46 0l2.1-2.09a2.45 2.45 0 0 1 3.47 0h0a2.45 2.45 0 0 1 0 3.46l-2.13 2.01A15 15 0 0 1 39 37v2H9v-2a15 15 0 0 1 2.84-8.77l-2.1-2.07a2.45 2.45 0 0 1 0-3.46h0a2.38 2.38 0 0 1 1.68-.7Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.68 31.17a2.75 2.75 0 0 0-2.77 2.76h0a2.77 2.77 0 0 0 2.77 2.77h0a2.77 2.77 0 0 0 0-5.53Zm12.64 0a2.77 2.77 0 1 0 2.77 2.76h0a2.75 2.75 0 0 0-2.77-2.76Z"/>`),
		g.Group(children),
	)
}