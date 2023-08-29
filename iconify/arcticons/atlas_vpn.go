package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtlasVpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.277 32.53A11.649 11.649 0 0 1 27.5 33.85h0l.003 11.647M3.305 29.857q-.028-.481-.028-.97a16.61 16.61 0 1 1 33.22 0v0c0 4.587.004 16.61.004 16.61"/><circle cx="16.501" cy="33.499" r="5.033" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".933"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24A21.5 21.5 0 1 0 24 45.5h21.5Z"/>`),
		g.Group(children),
	)
}