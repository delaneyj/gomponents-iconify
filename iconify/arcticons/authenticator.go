package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Authenticator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 27.95l19.5-10.43L24 8.13L4.5 17.52L24 27.95z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.37 20.66L4.5 23.48L24 33.91l19.5-10.43l-5.87-2.82"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.37 26.62L4.5 29.44L24 39.87l19.5-10.43l-5.87-2.82"/>`),
		g.Group(children),
	)
}