package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blabber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m24 42.5l-5.9-5.9h-6.3c-3.5 0-6.3-2.6-6.3-5.9V11.4c0-3.3 2.6-5.9 5.9-5.9h25.2c3.2 0 5.9 2.6 5.9 5.9c0 0 0 0 0 0v19.3c0 3.3-2.6 5.9-5.9 5.9h-6.7L24 42.5z"/><path d="M24 17.7c1.6 0 2.9 1.3 2.9 2.9s-1.3 2.9-2.9 2.9s-2.9-1.3-2.9-2.9s1.3-2.9 2.9-2.9zm-8.4 0c-1.6 0-2.9 1.3-2.9 2.9s1.3 2.9 2.9 2.9s2.9-1.3 2.9-2.9s-1.3-2.9-2.9-2.9zm16.8 0c1.6 0 2.9 1.3 2.9 2.9s-1.3 2.9-2.9 2.9s-2.9-1.3-2.9-2.9s1.3-2.9 2.9-2.9z"/></g>`),
		g.Group(children),
	)
}