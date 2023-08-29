package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cnn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 15.935v16.13l-12.56-16.13v16.13l-12.56-16.13v16.13h-5.815a8.065 8.065 0 0 1 0-16.13h2.829"/>`),
		g.Group(children),
	)
}