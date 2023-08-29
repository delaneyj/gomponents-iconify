package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTriangleRightCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M15.998 13L9 8.963v8.074L15.998 13Zm2.5.866a1 1 0 0 0 0-1.732L8.5 6.365a1 1 0 0 0-1.5.866V18.77a1 1 0 0 0 1.5.866l9.999-5.769Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTriangleRightCircleFilled0)"/></g>`),
		g.Group(children),
	)
}