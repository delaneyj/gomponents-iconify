package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phoneaccountabusedetector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.582 24.338v7.492m0-7.492a10.584 10.584 0 0 0-16.309-8.9m-2.751 2.567a10.536 10.536 0 0 0-2.104 6.333v13.895M9.67 41.981a3.748 3.748 0 0 1 3.748-3.748"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.67 41.981h28.66a3.748 3.748 0 0 0-3.748-3.748H13.418M7.91 9.418l32.344 32.343M5.5 24.52h4.691m27.618 0H42.5M24 10.71V6.02m9.764 8.735l3.317-3.317"/>`),
		g.Group(children),
	)
}