package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gcash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.824 33.23a16.122 16.122 0 0 0-.01-18.475m-1.806-2.149a16.114 16.114 0 1 0 0 22.788"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.461 16.923A9.18 9.18 0 1 0 29.794 24h-8.066m17.648 13.109a22.897 22.897 0 0 0-.015-26.24"/>`),
		g.Group(children),
	)
}