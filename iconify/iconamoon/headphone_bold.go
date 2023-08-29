package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2.5" d="M21 15v-3a9 9 0 1 0-18 0v3"/><path fill="currentColor" d="M3 15v-1.25H1.75V15H3Zm0 1.25h3v-2.5H3v2.5ZM5.75 16v3h2.5v-3h-2.5Zm-1.5 3v-4h-2.5v4h2.5Zm.75.75a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 5 22.25v-2.5Zm.75-.75a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 8.25 19h-2.5ZM6 16.25a.25.25 0 0 1-.25-.25h2.5A2.25 2.25 0 0 0 6 13.75v2.5ZM21 15h1.25v-1.25H21V15Zm-3 1.25h3v-2.5h-3v2.5ZM19.75 15v4h2.5v-4h-2.5Zm-1.5 4v-3h-2.5v3h2.5Zm.75.75a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 19 22.25v-2.5Zm.75-.75a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 22.25 19h-2.5ZM18 13.75A2.25 2.25 0 0 0 15.75 16h2.5a.25.25 0 0 1-.25.25v-2.5Z"/></g>`),
		g.Group(children),
	)
}