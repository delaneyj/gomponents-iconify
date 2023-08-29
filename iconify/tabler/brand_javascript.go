package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandJavascript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m20 4l-2 14.5l-6 2l-6-2L4 4z"/><path d="M7.5 8h3v8l-2-1m8-7H14a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5h1.423a.5.5 0 0 1 .495.57L15.5 15.5l-2 .5"/></g>`),
		g.Group(children),
	)
}