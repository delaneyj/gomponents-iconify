package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BancoAzteca(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.025 12.703c-7.55 1.594-16.035 4.767-19.91 11.042c-3.522 5.703-1.16 10.534 4.26 11.433c-8.492.725-12.07-1.901-11.868-7.216c.19-5.002 6.92-9.737 11.868-12.085c4.743-2.251 11.143-2.797 15.65-3.174Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.975 12.703c7.55 1.594 16.035 4.767 19.91 11.042c3.522 5.703 1.16 10.534-4.26 11.433c8.492.725 12.07-1.901 11.868-7.216c-.19-5.002-6.92-9.737-11.868-12.085c-4.743-2.251-11.143-2.797-15.65-3.174Z"/><circle cx="24.047" cy="27.115" r="6.422" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}