package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnreadBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M20.535 20.535C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465Z" opacity=".5"/><path d="M17.454 6.903a.75.75 0 0 1 .143 1.052l-6.858 9a.75.75 0 0 1-1.161.038l-3.143-3.6a.75.75 0 1 1 1.13-.986l2.538 2.907l6.3-8.268a.75.75 0 0 1 1.052-.143Z"/></g>`),
		g.Group(children),
	)
}