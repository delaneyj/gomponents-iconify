package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropletFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaDropletFill0"><g id="evaDropletFill1"><path id="evaDropletFill2" fill="currentColor" d="M12 21.1a7.4 7.4 0 0 1-5.28-2.28a7.73 7.73 0 0 1 .1-10.77l4.64-4.65a.94.94 0 0 1 .71-.3a1 1 0 0 1 .71.31l4.56 4.72a7.73 7.73 0 0 1-.09 10.77A7.33 7.33 0 0 1 12 21.1Z"/></g></g>`),
		g.Group(children),
	)
}