package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SecureVpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.646 43.297h12.708c2.963 0 5.7-1.58 7.18-4.145l6.355-11.006a8.291 8.291 0 0 0 0-8.292L37.535 8.848a8.291 8.291 0 0 0-7.18-4.145h-12.71a8.29 8.29 0 0 0-7.18 4.145L4.111 19.854a8.291 8.291 0 0 0 0 8.292l6.354 11.006a8.291 8.291 0 0 0 7.18 4.145Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="18.603" cy="24" r="4.936"/><path d="M34.332 24.018H23.539m7.176 2.953v-2.953"/></g>`),
		g.Group(children),
	)
}