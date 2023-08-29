package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMinusCircleFill0"><g id="evaMinusCircleFill1"><path id="evaMinusCircleFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm3 11H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2Z"/></g></g>`),
		g.Group(children),
	)
}