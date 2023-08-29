package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feClock0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feClock1" fill="currentColor"><path id="feClock2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm8-10a8 8 0 1 0-16 0a8 8 0 0 0 16 0Zm-4-1a1 1 0 0 1 0 2h-3c-1.1 0-2-.9-2-2V7a1 1 0 0 1 2 0v4h3Z"/></g></g>`),
		g.Group(children),
	)
}