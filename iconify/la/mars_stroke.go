package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarsStroke(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 4v2h7.563L20 10.563L17.719 8.28L16.28 9.72L18.563 12l-1.688 1.688C15.523 12.645 13.832 12 12 12c-4.406 0-8 3.594-8 8c0 4.406 3.594 8 8 8c4.406 0 8-3.594 8-8c0-1.832-.645-3.523-1.688-4.875L20 13.437l2.281 2.282l1.438-1.438L21.437 12L26 7.437V15h2V4zm-5 10c3.324 0 6 2.676 6 6s-2.676 6-6 6s-6-2.676-6-6s2.676-6 6-6z"/>`),
		g.Group(children),
	)
}