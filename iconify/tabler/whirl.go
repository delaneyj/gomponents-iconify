package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whirl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 12a2 2 0 1 0-4 0a2 2 0 0 0 4 0z"/><path d="M12 21c-3.314 0-6-2.462-6-5.5S8.686 10 12 10"/><path d="M21 12c0 3.314-2.462 6-5.5 6S10 15.314 10 12"/><path d="M12 14c3.314 0 6-2.462 6-5.5S15.314 3 12 3"/><path d="M14 12c0-3.314-2.462-6-5.5-6S3 8.686 3 12"/></g>`),
		g.Group(children),
	)
}