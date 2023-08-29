package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandBitbucket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.648 4a.64.64 0 0 0-.64.744l3.14 14.528c.07.417.43.724.852.728h10a.644.644 0 0 0 .642-.539l3.35-14.71a.641.641 0 0 0-.64-.744L3.648 4z"/><path d="M14 15h-4L9 9h6z"/></g>`),
		g.Group(children),
	)
}