package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clementineremote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.487 32.912c-.11 2.958-3.813 5.52-6.344 6.842h0c-1.795.821-4.284 1.401-7.285 1.573c-9.069.625-19.362-2.797-22.854-12.174h0c-1.325-3.033-1.995-5.836-2.291-8.255h0c-.363-2.867-.243-5.304.216-7.128h0c.505-3.038 3.249-6.272 5.9-7.012"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.487 32.912c-.516 4.834-13.332 6.057-19.204 4.31C14.208 35.575 8.928 26.363 7.85 17.392h0c-.943-12.42 4.83-14.571 11.385-4.082h0c1.598 2.07 4.285 6.432 5.928 7.773c4.449 2.852 18.83 6.402 18.324 11.83v-.001Z"/>`),
		g.Group(children),
	)
}