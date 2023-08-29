package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseOctagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.35 1.5l6.15 6.15v8.699l-6.15 6.15h-8.7L1.5 16.35v-8.7L7.65 1.5h8.7Zm-.83 2H8.48L3.5 8.479v7.041l4.98 4.98h7.04l4.98-4.98V8.48L15.52 3.5Zm1.076 5.318L13.414 12l3.182 3.181l-1.414 1.415L12 13.414l-3.182 3.182l-1.415-1.414L10.586 12L7.403 8.817l1.415-1.414L12 10.585l3.182-3.181l1.414 1.414Z"/>`),
		g.Group(children),
	)
}