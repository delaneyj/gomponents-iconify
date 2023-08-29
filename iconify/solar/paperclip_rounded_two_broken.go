package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperclipRoundedTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M19.565 16.13a8.25 8.25 0 0 0 0-11.706c-3.248-3.232-8.512-3.232-11.76 0m8.085 15.364a4.17 4.17 0 0 1-5.88 0a4.125 4.125 0 0 1 0-5.853l2.94-2.926M4.131 8.082a7.22 7.22 0 0 0 0 10.243"/>`),
		g.Group(children),
	)
}