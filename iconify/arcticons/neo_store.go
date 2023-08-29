package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeoStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.144 30.894A9.75 9.75 0 1 1 14.25 14.25c5.385 0 7.85 5.287 9.75 9.75c2.11 4.954 4.365 9.75 9.75 9.75a9.75 9.75 0 1 0-6.894-16.644"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.696 28.803a5 5 0 1 0-3.946-.022M7.75 24h13m-11-4h9m-9 8h9m18.903-12.935l1.562-1.738m-9.368 1.738l-1.562-1.738"/>`),
		g.Group(children),
	)
}