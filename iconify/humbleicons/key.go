package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="15.5" cy="8.5" r="1.5" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 16l5.18-5.652C10.033 9.875 10 9.523 10 9a5 5 0 1 1 5 5c-.523 0-.868-.01-1.342-.158L12 15.5h-2v2H8v2H5V16Z"/></g>`),
		g.Group(children),
	)
}