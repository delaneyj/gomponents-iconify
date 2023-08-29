package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluecoins(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.024 24.172a8.8 8.8 0 1 0-12.43-.198"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.557 28.249a8.8 8.8 0 1 0-6.009 10.883"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.28 28.223a8.8 8.8 0 1 0 6.183 10.785M26.4 19.848a8.8 8.8 0 1 0 6.11-10.826"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.802 19.798a8.8 8.8 0 1 0-6.286-10.724m13.021 15.282a8.8 8.8 0 1 0 12.42-.529"/>`),
		g.Group(children),
	)
}