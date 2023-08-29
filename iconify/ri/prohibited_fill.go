package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.743 18.329A9.958 9.958 0 0 0 22 12c0-5.523-4.477-10-10-10a9.959 9.959 0 0 0-6.329 2.257L19.743 18.33ZM4.257 5.67A9.959 9.959 0 0 0 2 12c0 5.523 4.477 10 10 10a9.958 9.958 0 0 0 6.329-2.257L4.257 5.67Z"/>`),
		g.Group(children),
	)
}