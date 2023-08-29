package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Antennapod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.59 42.72a21.5 21.5 0 1 0-21.18 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.33 36.53a15 15 0 1 0-16.74-.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30 30.14a8.6 8.6 0 1 0-12.08-.07"/><circle cx="24" cy="24" r="2.13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.27 26l-6.62 18.2M24.73 26l6.62 18.2"/>`),
		g.Group(children),
	)
}