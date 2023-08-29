package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvalFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2c3.972 0 7 4.542 7 10s-3.028 10-7 10c-3.9 0-6.89-4.379-6.997-9.703L5 12l.003-.297C5.11 6.38 8.1 2 12 2z"/></g>`),
		g.Group(children),
	)
}