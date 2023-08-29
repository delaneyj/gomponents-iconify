package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spanishdict(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="31.855" cy="31.943" r="4.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.253 4.5h-8.796l1.245 19.269h6.306L36.253 4.5z"/><circle cx="16.145" cy="16.057" r="4.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.747 43.5h8.796l-1.245-19.269h-6.306L11.747 43.5z"/>`),
		g.Group(children),
	)
}