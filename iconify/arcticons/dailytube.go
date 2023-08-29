package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dailytube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.12 5.5H8.38v37h12.74A18.5 18.5 0 0 0 39.62 24h0a18.5 18.5 0 0 0-18.5-18.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.073 24l-9.355-7.626v15.252L31.073 24zM8.381 5.5l13.337 10.874M8.381 42.5l13.337-10.874M31.073 24l2.56 2.087c2.7 1.964 3.836 7.39.586 10.976"/>`),
		g.Group(children),
	)
}