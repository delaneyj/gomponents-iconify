package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forkhub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="16.1" cy="38.92" r="4.58" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="16.1" cy="9.08" r="4.58" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.9" cy="24" r="4.58" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.1 34.34V13.66m0 19.01c0-5.4 5-8.67 11.22-8.67"/>`),
		g.Group(children),
	)
}