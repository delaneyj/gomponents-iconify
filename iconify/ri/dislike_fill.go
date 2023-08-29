package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DislikeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.808 1.394l18.384 18.384l-1.414 1.415l-3.746-3.747L12 21.486l-8.478-8.493a6 6 0 0 1 .033-8.023L1.394 2.808l1.414-1.414Zm17.435 3.364a6 6 0 0 1 .236 8.235l-1.635 1.636L7.26 3.046a5.99 5.99 0 0 1 4.741 1.483a5.998 5.998 0 0 1 8.242.229Z"/>`),
		g.Group(children),
	)
}