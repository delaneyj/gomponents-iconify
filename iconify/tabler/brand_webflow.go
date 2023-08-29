package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandWebflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 10s-1.376 3.606-1.5 4c-.046-.4-1.5-8-1.5-8c-2.627 0-3.766 1.562-4.5 3.5c0 0-1.843 4.593-2 5C7.487 14.132 7 10 7 10c-.15-2.371-2.211-3.98-4-3.98L5 19c2.745-.013 4.72-1.562 5.5-3.5c0 0 1.44-4.3 1.5-4.5c.013.18 1 8 1 8c2.758 0 4.694-1.626 5.5-3.5L22 6c-2.732 0-4.253 2.055-5 4z"/>`),
		g.Group(children),
	)
}