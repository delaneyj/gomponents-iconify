package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M20 12a8 8 0 1 1-16 0a8 8 0 0 1 16 0Z" opacity=".5"/><path d="M12 8.512a3.488 3.488 0 1 0 0 6.976a3.488 3.488 0 0 0 0-6.976ZM12.75 2a.75.75 0 0 0-1.5 0v2.035a8.102 8.102 0 0 1 1.5 0V2Zm7.215 10.75a8.085 8.085 0 0 0 0-1.5H22a.75.75 0 0 1 0 1.5h-2.035Zm-8.715 7.215a8.085 8.085 0 0 0 1.5 0V22a.75.75 0 0 1-1.5 0v-2.035ZM4.035 11.25a8.102 8.102 0 0 0 0 1.5H2a.75.75 0 0 1 0-1.5h2.035Z"/></g>`),
		g.Group(children),
	)
}