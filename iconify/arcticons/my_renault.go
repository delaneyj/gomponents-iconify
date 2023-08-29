package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyRenault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.75 4.5L10.458 24L20.75 43.5L31.042 24l-5.417-10.292"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.25 43.5L37.542 24L27.25 4.5L16.958 24l5.417 10.292"/>`),
		g.Group(children),
	)
}