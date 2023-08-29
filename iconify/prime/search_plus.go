package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.77 18.3a7.53 7.53 0 1 1 7.53-7.53a7.53 7.53 0 0 1-7.53 7.53Zm0-13.55a6 6 0 1 0 6 6a6 6 0 0 0-6-6Z"/><path fill="currentColor" d="M20 20.75a.74.74 0 0 1-.53-.22l-4.13-4.13a.75.75 0 0 1 1.06-1.06l4.13 4.13a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22ZM10.75 14a.76.76 0 0 1-.75-.75v-5a.75.75 0 0 1 1.5 0v5a.76.76 0 0 1-.75.75Z"/><path fill="currentColor" d="M13.25 11.5h-5a.75.75 0 0 1 0-1.5h5a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}