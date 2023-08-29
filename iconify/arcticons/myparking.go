package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myparking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.269 37.5h2.519m2.217 0h2.519m-7.255-3.375h2.519m2.217 0h2.519m-7.255-3.375h2.519m2.217 0h2.519m-7.255-3.375h2.519m2.217 0h2.519M14.269 24h2.519m2.217 0h2.519m-7.255-3.375h2.519m2.217 0h2.519m-7.255-3.375h2.519m2.217 0h2.519m6.952 3.375h2.519m2.217 0h2.519m-7.255-3.375h2.519m2.217 0h2.519m-21.462-3.375h2.519m2.217 0h2.519M14.269 10.5h2.519m2.217 0h2.519m2.217 16.875h2.518m2.217 0h2.519M23.741 24h2.518m2.217 0h2.519m-7.254-10.125h2.518m2.217 0h2.519M23.741 10.5h2.518m2.217 0h2.519m2.217 3.375h2.519M33.212 24h2.519"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}