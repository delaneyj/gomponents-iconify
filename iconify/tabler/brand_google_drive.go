package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGoogleDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 10L6 20l-3-5L9 5z"/><path d="M9 15h12l-3 5H6m9-5L9 5h6l6 10z"/></g>`),
		g.Group(children),
	)
}