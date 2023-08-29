package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBackspaceFill0"><g id="evaBackspaceFill1"><path id="evaBackspaceFill2" fill="currentColor" d="M20.14 4h-9.77a3 3 0 0 0-2 .78l-.1.11l-6 7.48a1 1 0 0 0 .11 1.37l6 5.48a3 3 0 0 0 2 .78h9.77A1.84 1.84 0 0 0 22 18.18V5.82A1.84 1.84 0 0 0 20.14 4Zm-3.43 9.29a1 1 0 0 1 0 1.42a1 1 0 0 1-1.42 0L14 13.41l-1.29 1.3a1 1 0 0 1-1.42 0a1 1 0 0 1 0-1.42l1.3-1.29l-1.3-1.29a1 1 0 0 1 1.42-1.42l1.29 1.3l1.29-1.3a1 1 0 0 1 1.42 1.42L15.41 12Z"/></g></g>`),
		g.Group(children),
	)
}