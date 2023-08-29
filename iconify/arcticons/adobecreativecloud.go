package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobecreativecloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2Zm14.21 29.23c-3.3.12-6.06.32-8.84-1.51a8.86 8.86 0 0 1 6.43-16.07c3.24.56 4.61 2.41 6.19 4c.44.43.83.84 1.26 1.28a2.46 2.46 0 0 1 1 1.55m-6.19-8.12c.14-.46 2.3-1.56 2.65-1.71A10.64 10.64 0 0 1 39 23.88a10.59 10.59 0 0 1-.64 3.76a9.8 9.8 0 0 1-1.7 3.08a10.52 10.52 0 0 1-10.06 3.87c-4-.67-6-3.16-8.48-5.7c-.5-.51-2.12-2-2.27-2.58"/>`),
		g.Group(children),
	)
}