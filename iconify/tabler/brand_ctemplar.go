package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCtemplar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.04 14.831L10.5 10.5m2.055 10.32c4.55-3.456 7.582-8.639 8.426-14.405a1.668 1.668 0 0 0-.934-1.767A19.647 19.647 0 0 0 12 3a19.647 19.647 0 0 0-8.047 1.647a1.668 1.668 0 0 0-.934 1.767c.844 5.766 3.875 10.95 8.426 14.406a.948.948 0 0 0 1.11 0z"/><path d="M20 5c-2 0-4.37 3.304-8 6.644C8.37 8.304 6 5 4 5m13.738 10L13.5 10.5"/></g>`),
		g.Group(children),
	)
}