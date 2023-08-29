package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OtpAuthenticator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.83 26.342a7.835 7.835 0 0 1-9.498-2.61a7.276 7.276 0 0 1 .982-9.516a7.878 7.878 0 0 1 9.845-.736a7.308 7.308 0 0 1 2.479 9.257m-3.702 3.607l4.082 4.083h2.295v2.19h2.241v2.248l1.04 1.041H36v-3.278l-9.503-9.503"/><circle cx="17.498" cy="17.337" r="2.288" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}