package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.25 4.5h12.12c4.18 0 7.44.88 9.24 2.94c1.64 1.87 2.13 3.94 1.65 7a6.09 6.09 0 0 1 1.74 1.3C40.53 17.5 40.83 20 40.25 23c-1.4 7.18-6.18 9.67-12.29 9.67H27A1.52 1.52 0 0 0 25.51 34l-.07.41l-1.19 7.49l-.06.32a1.52 1.52 0 0 1-1.5 1.28h-6.31a1.41 1.41 0 0 1-1.52-1.5c.84-5.36 1.65-10.39 2.49-15.74a5.25 5.25 0 0 1 1.49-.06h4.24c7 0 12.46-2.83 14.06-11l.12-.71a9.21 9.21 0 0 0-1.26-.64l-.36-.11l-.76-.19c-.26-.06-.52-.11-.8-.15a19.18 19.18 0 0 0-3.13-.23H20.71a1.45 1.45 0 0 0-.66.15a1.52 1.52 0 0 0-.84 1.13q-1.93 12.32-3.89 24.61v.11h-6.8a1 1 0 0 1-1-1.2l5-32a1.74 1.74 0 0 1 1.73-1.47Z"/>`),
		g.Group(children),
	)
}