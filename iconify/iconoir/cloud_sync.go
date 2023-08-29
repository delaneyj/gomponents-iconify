package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607m3.58 1.88l1.768 1.768a4 4 0 0 0 5.657 0l.354-.353"/><path d="m7.934 21.962l-.353-2.475l2.474.354l-2.12 2.121Zm8.364-5.06l-1.768-1.768a4 4 0 0 0-5.657 0l-.353.353"/><path d="m15.944 14.427l.354 2.475l-2.475-.354l2.121-2.121Z"/></g>`),
		g.Group(children),
	)
}