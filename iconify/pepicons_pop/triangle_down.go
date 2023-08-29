package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 12.998L14.037 6H5.963L10 12.998Zm-.866 2.5a1 1 0 0 0 1.732 0L16.635 5.5a1 1 0 0 0-.866-1.5H4.23a1 1 0 0 0-.866 1.5l5.769 9.999Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}