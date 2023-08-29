package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oculus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.054 36.446H16.946a12.446 12.446 0 0 1 0-24.892h14.108a12.446 12.446 0 1 1 0 24.892Zm-14.108-16.46a4.013 4.013 0 1 0 0 8.027h14.108a4.013 4.013 0 1 0 0-8.026Z"/>`),
		g.Group(children),
	)
}