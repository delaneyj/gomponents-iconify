package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42h0l7.18-22.1H16.82L24 42ZM6.76 19.86h0l-2.19 6.71a1.5 1.5 0 0 0 .54 1.67L24 42L6.76 19.86Zm0 0h10.06L12.49 6.55a.74.74 0 0 0-1.41 0L6.76 19.86Zm34.49 0h0l2.18 6.71a1.5 1.5 0 0 1-.54 1.67L24 42l17.25-22.1Zm0 0H31.18l4.33-13.31a.74.74 0 0 1 1.41 0l4.33 13.31Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 41.96l7.18-22.1h10.07L24 41.96zm0 0L6.76 19.86h10.06L24 41.96z"/>`),
		g.Group(children),
	)
}