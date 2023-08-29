package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backgrounds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="21.064" height="32.17" x="10.037" y="11.33" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.561" ry="2.561"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.452 11.33v-.854a2.561 2.561 0 0 1 2.561-2.561h15.942a2.561 2.561 0 0 1 2.561 2.561v27.048a2.561 2.561 0 0 1-2.561 2.561H31.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.899 7.915V7.06A2.561 2.561 0 0 1 19.46 4.5h15.942a2.561 2.561 0 0 1 2.56 2.561V34.11a2.561 2.561 0 0 1-2.56 2.561h-.886"/>`),
		g.Group(children),
	)
}