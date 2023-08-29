package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDuoLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.585 18.01L5.575 12l6.01-6.01L13 7.404l-4.6 4.6l4.6 4.6l-1.414 1.406h-.001Zm5.425 0L10.999 12l6.011-6.01l1.414 1.414l-4.6 4.6l4.6 4.6l-1.413 1.406h-.001Z"/>`),
		g.Group(children),
	)
}