package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piavpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.977 25.28a3.261 3.261 0 0 0-1.634 6.074l-1 4.64h5.312l-1-4.64a3.266 3.266 0 0 0 1.61-2.806v-.038a3.266 3.266 0 0 0-3.288-3.23Z"/><rect width="31.231" height="25.727" x="8.384" y="17.773" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.352 17.773V15.16a10.66 10.66 0 0 1 21.319 0v2.613"/><path fill="currentColor" d="M21.108 10.984a.75.75 0 1 0-.75.75a.75.75 0 0 0 .75-.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.449 14.607a6.108 6.108 0 0 0 2.562.655a6.128 6.128 0 0 0 2.563-.655"/><path fill="currentColor" d="M27.665 10.234a.75.75 0 1 0 .75.75a.75.75 0 0 0-.75-.75Z"/>`),
		g.Group(children),
	)
}