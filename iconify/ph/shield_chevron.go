package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldChevron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 40H48a16 16 0 0 0-16 16v58.77c0 89.61 75.82 119.34 91 124.39a15.53 15.53 0 0 0 10 0c15.2-5.05 91-34.78 91-124.39V56a16 16 0 0 0-16-16Zm-80 184c-9.26-3.08-43.29-16.32-63.87-49.5L128 129.76l63.87 44.71C171.31 207.61 137.34 220.85 128 224Zm80-109.18c0 17.64-3.36 32.63-8.72 45.34l-66.69-46.68a8 8 0 0 0-9.18 0l-66.69 46.65c-5.36-12.71-8.72-27.7-8.72-45.34V56h160Z"/>`),
		g.Group(children),
	)
}