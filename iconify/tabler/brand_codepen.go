package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCodepen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 15l9 6l9-6l-9-6l-9 6"/><path d="m3 9l9 6l9-6l-9-6l-9 6m0 0v6m18-6v6M12 3v6m0 6v6"/></g>`),
		g.Group(children),
	)
}