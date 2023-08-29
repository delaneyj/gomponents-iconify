package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConsorsbankSecureplus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="33.231" height="25.727" x="7.384" y="17.773" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.316"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.352 17.773V15.16a10.66 10.66 0 0 1 21.32 0v2.613M24 34.636v-8"/>`),
		g.Group(children),
	)
}