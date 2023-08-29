package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMessageCircleFill0"><g id="evaMessageCircleFill1"><path id="evaMessageCircleFill2" fill="currentColor" d="M19.07 4.93a10 10 0 0 0-16.28 11a1.06 1.06 0 0 1 .09.64L2 20.8a1 1 0 0 0 .27.91A1 1 0 0 0 3 22h.2l4.28-.86a1.26 1.26 0 0 1 .64.09a10 10 0 0 0 11-16.28ZM8 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`),
		g.Group(children),
	)
}