package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkCircleTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCheckmarkCircle2Outline0"><g id="evaCheckmarkCircle2Outline1"><g id="evaCheckmarkCircle2Outline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="m14.7 8.39l-3.78 5l-1.63-2.11a1 1 0 0 0-1.58 1.23l2.43 3.11a1 1 0 0 0 .79.38a1 1 0 0 0 .79-.39l4.57-6a1 1 0 1 0-1.6-1.22Z"/></g></g></g>`),
		g.Group(children),
	)
}