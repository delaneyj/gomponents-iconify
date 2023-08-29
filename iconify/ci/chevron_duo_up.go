package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDuoUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.403 18.425L5.99 17.01L12 10.999l6.01 6.011l-1.413 1.413l-4.6-4.6l-4.6 4.6l.006.002Zm0-5.424L5.99 11.585L12 5.575l6.01 6.01l-1.413 1.414l-4.6-4.6l-4.6 4.6l.006.002Z"/>`),
		g.Group(children),
	)
}