package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandApplePodcast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18.364 18.364a9 9 0 1 0-12.728 0"/><path d="M11.766 22h.468a2 2 0 0 0 1.985-1.752l.5-4A2 2 0 0 0 12.734 14h-1.468a2 2 0 0 0-1.985 2.248l.5 4A2 2 0 0 0 11.766 22zM10 9a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/></g>`),
		g.Group(children),
	)
}