package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m26.24 16.373l-9.14-9.14c-2.661-2.661-7.035-2.603-9.768.131c-2.734 2.734-2.793 7.107-.131 9.768l7.935 7.936m17.767-2.065l7.935 7.935c2.661 2.662 2.603 7.035-.13 9.769c-2.735 2.734-7.108 2.792-9.77.13l-9.14-9.14"/><path d="M26.11 26.142c2.733-2.734 2.792-7.108.13-9.769m-4.441 5.425c-2.734 2.734-2.792 7.108-.131 9.769"/></g>`),
		g.Group(children),
	)
}