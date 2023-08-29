package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M19 5h-9v14h9a2 2 0 0 0 2-2v-3a2 2 0 1 1 0-4V7a2 2 0 0 0-2-2Z" opacity=".16"/><path d="M3 10H2a1 1 0 0 0 1 1v-1Zm0 4v-1a1 1 0 0 0-1 1h1Zm18-4v1a1 1 0 0 0 1-1h-1Zm0 4h1a1 1 0 0 0-1-1v1ZM5 6h5V4H5v2Zm5 0h9V4h-9v2Zm9 12h-9v2h9v-2Zm-9 0H5v2h5v-2ZM9 5v14h2V5H9Zm-5.293 6.293a1 1 0 0 1 0 1.414l1.414 1.414a3 3 0 0 0 0-4.242l-1.414 1.414Zm16.586 1.414a1 1 0 0 1 0-1.414l-1.414-1.414a3 3 0 0 0 0 4.242l1.414-1.414ZM3 11c.257 0 .512.097.707.293l1.414-1.414A2.994 2.994 0 0 0 3 9v2Zm1-1V7H2v3h2Zm0 7v-3H2v3h2Zm-.293-4.293A.994.994 0 0 1 3 13v2c.766 0 1.536-.293 2.121-.879l-1.414-1.414Zm16.586-1.414A.994.994 0 0 1 21 11V9c-.766 0-1.536.293-2.121.879l1.414 1.414ZM20 7v3h2V7h-2Zm0 7v3h2v-3h-2Zm1-1a.994.994 0 0 1-.707-.293l-1.414 1.414A2.994 2.994 0 0 0 21 15v-2ZM5 18a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3v-2Zm14 2a3 3 0 0 0 3-3h-2a1 1 0 0 1-1 1v2Zm0-14a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM5 4a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V4Z"/></g>`),
		g.Group(children),
	)
}