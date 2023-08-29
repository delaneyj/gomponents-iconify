package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYangFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M17 3.34a10 10 0 1 1-14.995 8.984L2 12l.005-.324A10 10 0 0 1 17 3.34zM8 5.072A8 8 0 0 0 12 20l.2-.005a4 4 0 0 0 0-7.99L12 12a4 4 0 0 1-.2-7.995L12 4a7.995 7.995 0 0 0-4 1.072zM12 6.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3z"/><path fill="currentColor" d="M12 14.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3z"/></g>`),
		g.Group(children),
	)
}