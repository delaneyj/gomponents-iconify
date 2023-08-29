package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDingtalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0z"/><path d="m8 7.5l7.02 2.632a1 1 0 0 1 .567 1.33L14.5 14H16l-5 4l1-4c-3.1.03-3.114-3.139-4-6.5z"/></g>`),
		g.Group(children),
	)
}