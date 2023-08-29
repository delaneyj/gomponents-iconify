package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightFilledCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTriangleRightFilledCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M18.499 12.134a1 1 0 0 1 0 1.732l-10 5.769A1 1 0 0 1 7 18.769V7.23a1 1 0 0 1 1.5-.866l9.999 5.769Z"/><path d="M8.5 19.635a1 1 0 0 1-1.5-.866V7.23a1 1 0 0 1 1.5-.866l9.999 5.769a1 1 0 0 1 0 1.732l-10 5.769ZM13.997 13L10 10.694v4.612L13.997 13Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTriangleRightFilledCircleFilled0)"/></g>`),
		g.Group(children),
	)
}