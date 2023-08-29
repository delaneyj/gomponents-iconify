package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandOauth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12a10 10 0 1 0 20 0a10 10 0 1 0-20 0"/><path d="M12.556 6c.65 0 1.235.373 1.508.947l2.839 7.848a1.646 1.646 0 0 1-1.01 2.108a1.673 1.673 0 0 1-2.068-.851L13.365 15h-2.73l-.398.905A1.67 1.67 0 0 1 8.26 16.95l-.153-.047a1.647 1.647 0 0 1-1.056-1.956l2.824-7.852a1.664 1.664 0 0 1 1.409-1.087L12.556 6z"/></g>`),
		g.Group(children),
	)
}