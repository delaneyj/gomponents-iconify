package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTriangleDownCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M13 15.998L17.037 9H8.963L13 15.998Zm-.866 2.5a1 1 0 0 0 1.732 0L19.635 8.5a1 1 0 0 0-.866-1.5H7.23a1 1 0 0 0-.866 1.5l5.769 9.999Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTriangleDownCircleFilled0)"/></g>`),
		g.Group(children),
	)
}