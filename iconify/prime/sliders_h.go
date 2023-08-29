package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 8.25h-7a.75.75 0 0 1 0-1.5h7a.75.75 0 0 1 0 1.5Zm-9 0H5a.75.75 0 0 1 0-1.5h5a.75.75 0 0 1 0 1.5Z"/><path fill="currentColor" d="M10 9.75A.76.76 0 0 1 9.25 9V6a.75.75 0 0 1 1.5 0v3a.76.76 0 0 1-.75.75Zm9 7.5h-7a.75.75 0 0 1 0-1.5h7a.75.75 0 0 1 0 1.5Zm-9 0H5a.75.75 0 0 1 0-1.5h5a.75.75 0 0 1 0 1.5Z"/><path fill="currentColor" d="M10 18.75a.76.76 0 0 1-.75-.75v-3a.75.75 0 0 1 1.5 0v3a.76.76 0 0 1-.75.75Zm9-6h-3a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5Zm-5 0H5a.75.75 0 0 1 0-1.5h9a.75.75 0 0 1 0 1.5Z"/><path fill="currentColor" d="M14 14.25a.76.76 0 0 1-.75-.75v-3a.75.75 0 0 1 1.5 0v3a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}