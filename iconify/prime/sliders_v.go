package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 19.75a.76.76 0 0 1-.75-.75v-7a.75.75 0 0 1 1.5 0v7a.76.76 0 0 1-.75.75Zm0-9a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v5a.76.76 0 0 1-.75.75Z"/><path fill="currentColor" d="M18 10.75h-3a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5Zm-10.5 9a.76.76 0 0 1-.75-.75v-7a.75.75 0 0 1 1.5 0v7a.76.76 0 0 1-.75.75Zm0-9a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v5a.76.76 0 0 1-.75.75Z"/><path fill="currentColor" d="M9 10.75H6a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5Zm3 9a.76.76 0 0 1-.75-.75v-3a.75.75 0 0 1 1.5 0v3a.76.76 0 0 1-.75.75Zm0-5a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v9a.76.76 0 0 1-.75.75Z"/><path fill="currentColor" d="M13.5 14.75h-3a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}