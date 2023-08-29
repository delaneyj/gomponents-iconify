package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSharik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.281 16.606A8.968 8.968 0 0 1 5.644 5.629a9.033 9.033 0 0 1 11.011-1.346C15.071 8.975 14.24 11.243 12 13c-1.584 1.242-3.836 2.24-7.719 3.606zM20.616 9.3c2.113 7.59-4.892 13.361-11.302 11.264c1.931-3.1 3.235-4.606 4.686-6.065c1.705-1.715 3.591-3.23 6.616-5.199z"/>`),
		g.Group(children),
	)
}