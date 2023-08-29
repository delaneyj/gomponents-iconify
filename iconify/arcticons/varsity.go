package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Varsity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.4 4.5a2 2 0 0 1 2 2v19.1c0 6.88-10.398 17.9-16.4 17.9c-6.002 0-16.4-11.02-16.4-17.9V6.5a2 2 0 0 1 2-2h28.8Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.4 12.6L24 35.4l-7.4-22.8"/>`),
		g.Group(children),
	)
}