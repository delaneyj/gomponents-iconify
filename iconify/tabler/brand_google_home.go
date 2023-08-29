package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGoogleHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.072 21H4.928A1.928 1.928 0 0 1 3 19.072v-6.857c0-.512.203-1 .566-1.365l7.07-7.063a1.928 1.928 0 0 1 2.727 0l7.071 7.063c.363.362.566.853.566 1.365v6.857A1.928 1.928 0 0 1 19.072 21z"/><path d="M7 13v4h10v-4l-5-5m2.8-2.8L3 17m4 0v4m10-4v4"/></g>`),
		g.Group(children),
	)
}