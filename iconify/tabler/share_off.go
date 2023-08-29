package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0m12-6a3 3 0 1 0 6 0a3 3 0 1 0-6 0m.861 9.896a3 3 0 0 0 4.265 4.22m.578-3.417a3.012 3.012 0 0 0-1.507-1.45M8.7 10.7l1.336-.688M12.66 8.66L15.3 7.3m-6.6 6l6.6 3.4M3 3l18 18"/>`),
		g.Group(children),
	)
}