package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 37.382h37V42.5h-37zm37-22.817L24 5.5L5.5 14.565v4.683h37v-4.683zm-37 0h11.655m13.69 0H42.5"/><circle cx="24" cy="12.615" r="2.696" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.162 19.248v18.134M13.95 19.248v18.134M9.739 19.248v18.134m20.099 0V19.248m4.212 18.134V19.248m4.211 18.134V19.248"/>`),
		g.Group(children),
	)
}