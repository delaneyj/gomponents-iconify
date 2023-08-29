package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.638 5.634a9 9 0 1 0 12.723 12.733m1.686-2.332A9 9 0 0 0 7.954 3.958"/><path d="M10.5 10.5a2.187 2.187 0 0 0-2.914.116a1.928 1.928 0 0 0 0 2.768a2.188 2.188 0 0 0 2.914.116m6-3a2.194 2.194 0 0 0-2.309-.302M3 3l18 18"/></g>`),
		g.Group(children),
	)
}