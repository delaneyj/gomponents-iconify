package feather

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 18a5 5 0 0 0-10 0m5-16v7m-7.78 1.22l1.42 1.42M1 18h2m18 0h2m-4.64-6.36l1.42-1.42M23 22H1M8 6l4-4l4 4"/>`),
		g.Group(children),
	)
}