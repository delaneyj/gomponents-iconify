package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm.5 1.03c4.56.252 8.215 3.923 8.46 8.486H12.5V3.03zM12 21c-4.963 0-9-4.037-9-9c0-4.794 3.77-8.713 8.5-8.975V12a.5.5 0 0 0 .067.25l4.488 7.774A8.933 8.933 0 0 1 12 21zm4.917-1.482l-4.042-7.002h8.076a9.002 9.002 0 0 1-4.034 7.002z"/>`),
		g.Group(children),
	)
}