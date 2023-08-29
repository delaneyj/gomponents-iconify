package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lemon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.45 8.74A2.23 2.23 0 0 1 21.64 7a3.51 3.51 0 0 0 .24-2.47a3.55 3.55 0 0 0-2.45-2.45a3.51 3.51 0 0 0-2.43.28a2.23 2.23 0 0 1-1.7.19a10.07 10.07 0 0 0-6.53 0a9.87 9.87 0 0 0-6.23 6.18a10.07 10.07 0 0 0 0 6.53A2.23 2.23 0 0 1 2.36 17a3.51 3.51 0 0 0-.24 2.47a3.55 3.55 0 0 0 2.45 2.45A3.51 3.51 0 0 0 7 21.64a2.23 2.23 0 0 1 1.7-.19A9.83 9.83 0 0 0 12 22a10.33 10.33 0 0 0 3.27-.54a9.87 9.87 0 0 0 6.19-6.19a10.07 10.07 0 0 0-.01-6.53zM12 7a5 5 0 0 0-5 5H5a7 7 0 0 1 7-7z"/>`),
		g.Group(children),
	)
}