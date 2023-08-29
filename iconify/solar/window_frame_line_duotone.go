package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowFrameLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z"/><path fill="currentColor" d="M7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill="currentColor" d="M2 8.75a.75.75 0 0 0 0 1.5v-1.5Zm20 1.5a.75.75 0 0 0 0-1.5v1.5ZM8.25 21a.75.75 0 0 0 1.5 0h-1.5Zm1.5-11a.75.75 0 0 0-1.5 0h1.5ZM2 10.25h20v-1.5H2v1.5ZM9.75 21V10h-1.5v11h1.5Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}