package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnavailableFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.503 5.382a8 8 0 0 1 11.114 11.114L7.504 5.383Zm-2.12 2.121a8 8 0 0 0 11.114 11.114L5.382 7.504ZM12 2a9.972 9.972 0 0 0-7.071 2.929A9.972 9.972 0 0 0 2 12c0 5.523 4.477 10 10 10a9.972 9.972 0 0 0 7.071-2.929A9.972 9.972 0 0 0 22 12c0-5.523-4.477-10-10-10Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}