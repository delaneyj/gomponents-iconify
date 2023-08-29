package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMinimalisticSquareThreeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z" opacity=".5"/><path fill-rule="evenodd" d="M11.663 11.25a3.251 3.251 0 1 0 0 1.5h3.087v.75a.75.75 0 0 0 1.5 0v-.75H17a.25.25 0 0 1 .25.25v1a.75.75 0 0 0 1.5 0v-1A1.75 1.75 0 0 0 17 11.25h-5.337Zm-3.163-1a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}