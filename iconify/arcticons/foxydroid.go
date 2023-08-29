package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foxydroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 16.17l-9.24 5.33l-9.24 5.34V5.5l9.24 5.33L24 16.17zm0 0l9.24-5.34l9.24-5.33v21.34l-9.24-5.34L24 16.17z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.48 26.84L24 42.5V16.17l9.24 5.33l9.24 5.34zm-36.96 0l9.24-5.34L24 16.17V42.5L5.52 26.84z"/>`),
		g.Group(children),
	)
}