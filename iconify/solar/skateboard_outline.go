package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkateboardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m2.624 6.584l.813 1.219A3.25 3.25 0 0 0 6.14 9.25h11.718a3.25 3.25 0 0 0 2.704-1.447l.813-1.219l1.248.832l-.813 1.219a4.75 4.75 0 0 1-3.952 2.115H6.141a4.75 4.75 0 0 1-3.952-2.115l-.813-1.219l1.248-.832ZM7 13.75a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5ZM4.25 15a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0ZM17 13.75a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5ZM14.25 15a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}