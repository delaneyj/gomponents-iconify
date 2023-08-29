package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Preview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h37v37h-37zm0 24.7h37m-37-12.4h37M30.2 5.5v37m-12.4-37v37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.8 19.5c-.4 0-.7.3-.7.7s.3.7.7.7s.7-.3.7-.7h0c0-.4-.3-.7-.7-.7ZM24 21.2c-1.5 0-2.8 1.2-2.8 2.8s1.2 2.8 2.8 2.8c1.5 0 2.8-1.2 2.8-2.8h0c0-1.5-1.3-2.8-2.8-2.8Z"/>`),
		g.Group(children),
	)
}