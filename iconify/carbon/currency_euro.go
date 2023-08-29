package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyEuro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 26c-3.616 0-6.333-2.297-7.446-6H19v-2H9.132A15.278 15.278 0 0 1 9 16c0-.33.01-.664.026-1H19v-2H9.237C9.845 9.352 11.81 6 17 6c3.853 0 5.532 1.647 7.128 4.49l1.744-.98C24.265 6.649 22.078 4 17 4C10.645 4 7 8.374 7 16c0 7.065 4.112 12 10 12c5.078 0 7.265-2.648 8.872-5.51l-1.744-.98C22.532 24.354 20.853 26 17 26Z"/>`),
		g.Group(children),
	)
}