package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NethunterKex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.5 19.6a8.413 8.413 0 0 0-1.3-1.2c-.5-.3-.3-.9-.3-1.1s-3.6-2.2-4.7-2.2s-3.7.6-3.7 2.8s2.1 3 4.2 3.5c2.4.6 4.8 1 6.1 3a8.833 8.833 0 0 1 1 2.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.3 21.3s6 2.2 5.4 7.8m-5.1-13.9a3.26 3.26 0 0 1-2.2-3.3c0-.9-12.2-3.5-12.2-3.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23 12.4a37.322 37.322 0 0 0-11.2-.9m11.3 2.2s-5.5-.5-9.3.9m-1.3 22.7a4.232 4.232 0 0 0-1.6.2l-5 4.8c-.2.2.5.4 1.6.4h33.1c1.1 0 1.8-.2 1.6-.4l-5-4.8a4.232 4.232 0 0 0-1.6-.2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.3 31.5h15.4a3.598 3.598 0 0 1-.9 2.2a2.94 2.94 0 0 1-2.2.9h-9.3a3.097 3.097 0 0 1-3-3.1Zm0 0h24.2a1.923 1.923 0 0 0 1.9-2v-22a2.006 2.006 0 0 0-2-2h-33a2.006 2.006 0 0 0-2 2v22a2.006 2.006 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}