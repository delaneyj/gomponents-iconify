package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareArrowLeftDownBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.464 20.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535Z" opacity=".5"/><path d="M8.42 14.828c0 .414.337.75.75.75h4.244a.75.75 0 1 0 0-1.5h-2.432L15.358 9.7a.75.75 0 1 0-1.06-1.06L9.92 13.017v-2.432a.75.75 0 0 0-1.5 0v4.243Z"/></g>`),
		g.Group(children),
	)
}