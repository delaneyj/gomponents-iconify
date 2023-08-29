package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disappointed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDisappointed0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDisappointed1" fill="currentColor"><path id="feDisappointed2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm4-3h-1.339s-.417-2.667-2.661-2.667S9.333 17 9.333 17H8a4 4 0 1 1 8 0Zm-2.234-7.808L13 9V8l4 1v1l-1.237-.31a1.5 1.5 0 1 1-1.997-.5Zm-5.53.499L7 10V9l4-1v1l-.766.192a1.5 1.5 0 1 1-1.997.499Z"/></g></g>`),
		g.Group(children),
	)
}