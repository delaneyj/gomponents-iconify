package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.517 35.415H6.403a.903.903 0 0 1-.903-.903V21.008c0-.498.404-.902.903-.902h35.194c.499 0 .903.404.903.902v13.504a.903.903 0 0 1-.903.903h-8.114"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.517 31.437h18.966v10.13H14.517zm18.966-11.331H14.517V6.433h14.966l4 4v9.673Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.872 6.433v3.611h3.611"/><circle cx="39.852" cy="22.898" r=".75" fill="currentColor"/><circle cx="36.835" cy="22.898" r=".75" fill="currentColor"/><circle cx="33.819" cy="22.898" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}