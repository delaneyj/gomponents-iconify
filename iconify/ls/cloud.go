package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M343 620H171C77 620 0 543 0 449c0-71 43-130 103-156c0-5-1-10-1-15c0-56 47-102 103-102c34 0 64 17 83 42c31-84 112-145 207-145c122 0 220 99 222 220c60 27 102 86 102 156c0 94-77 171-171 171H481V520h81c12 0 16-8 7-17L428 352c-4-4-10-7-16-7s-13 3-17 7L254 503c-8 9-6 17 7 17h82v100z"/>`),
		g.Group(children),
	)
}