package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KulAuthenticator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.393 26.472L29.01 41.21l-.732 1.29l-5-2.515l-4.047-14.158"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.393 26.472c-6.605.37-12.467-3.46-11.261-11.166"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.675 19.276a7.41 7.41 0 1 1 11.133-5.333"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.337 19.411l12.554 11.947l.05 2.497l-5.634.03l-10.95-10.738m-.013.005a7.112 7.112 0 0 1-9.601-6.661h0a7.111 7.111 0 0 1 7.11-7.111h0a7.111 7.111 0 0 1 7.112 7.11h0a7.111 7.111 0 0 1-.623 2.91"/>`),
		g.Group(children),
	)
}