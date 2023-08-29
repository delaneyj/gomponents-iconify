package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Donut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDonut0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDonut1" fill="currentColor"><path id="feDonut2" d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 6C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-18h2v2h-2V4Zm6 5h2v2h-2V9ZM5 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm2 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-4 2h2v2h-2v-2Zm-4 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-5-4h2v2H6v-2ZM5 8h2v2H5V8Z"/></g></g>`),
		g.Group(children),
	)
}