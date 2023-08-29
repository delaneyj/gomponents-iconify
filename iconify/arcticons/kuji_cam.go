package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KujiCam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5c-1.1 0-2 .9-2 2h0v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2h0v-33c0-1.1-.9-2-2-2h-33ZM36 20v8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32 20v6c0 1.1-.9 2-2 2h0c-1.1 0-2-.9-2-2v-.7M20 20v5.3c0 1.5 1.2 2.7 2.6 2.7s2.6-1.2 2.6-2.7V20M12 20v8m0-2.8l4.2-5.2m0 8L13 24M5.5 11.9h37m-37 24.2h37"/>`),
		g.Group(children),
	)
}