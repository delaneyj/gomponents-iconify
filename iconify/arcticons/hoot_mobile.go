package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HootMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 36.464a24.316 24.316 0 0 0 9.125-8.114c4.136-5.945 8.838-11.32 15.863-14.422a16.497 16.497 0 0 0 4.278-2.428s-3.664 5.415-5.325 6.234c0 0 8.641-5.569 15.059-4.024c0 0-7.299 2.12-12.209 6.792a40.896 40.896 0 0 0 8.921-1.344c.03.064-.7 1.89-3.461 2.73c-3.268.994-6.778 1.794-7.796 2.947c0 0 5.092-.724 7.15.403c-5.19 1.524-11.046 2.986-18.916 7.559C9.742 37.123 4.5 36.464 4.5 36.464Z"/>`),
		g.Group(children),
	)
}