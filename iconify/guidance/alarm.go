package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m20.5 23l-2.664-3.148M3.5 23l2.664-3.148M22.912 8A12.027 12.027 0 0 0 15 1.378M1.088 8A12.027 12.027 0 0 1 9 1.378M12 6s.5 3.5 0 7c2.75 1.5 5 4 5 4M6.164 19.852a9 9 0 1 1 11.672 0m-11.672 0A8.964 8.964 0 0 0 12 22a8.965 8.965 0 0 0 5.836-2.148"/>`),
		g.Group(children),
	)
}