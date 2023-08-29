package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPauseCircleFill0"><g id="evaPauseCircleFill1"><path id="evaPauseCircleFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm-2 13a1 1 0 0 1-2 0V9a1 1 0 0 1 2 0Zm6 0a1 1 0 0 1-2 0V9a1 1 0 0 1 2 0Z"/></g></g>`),
		g.Group(children),
	)
}