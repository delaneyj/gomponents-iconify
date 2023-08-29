package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M21 15v-3a9 9 0 1 0-18 0v3"/><path fill="currentColor" d="M3 15v-.75h-.75V15H3Zm0 .75h3v-1.5H3v1.5Zm3.25.25v3h1.5v-3h-1.5Zm-2.5 3v-4h-1.5v4h1.5ZM5 20.25c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 5 21.75v-1.5ZM6.25 19c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 7.75 19h-1.5ZM6 15.75a.25.25 0 0 1 .25.25h1.5A1.75 1.75 0 0 0 6 14.25v1.5ZM21 15h.75v-.75H21V15Zm-3 .75h3v-1.5h-3v1.5Zm2.25-.75v4h1.5v-4h-1.5Zm-2.5 4v-3h-1.5v3h1.5ZM19 20.25c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 19 21.75v-1.5ZM20.25 19c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 21.75 19h-1.5ZM18 14.25A1.75 1.75 0 0 0 16.25 16h1.5a.25.25 0 0 1 .25-.25v-1.5Z"/></g>`),
		g.Group(children),
	)
}