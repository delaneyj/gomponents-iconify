package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalicTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.75 4h8.504a.75.75 0 0 1 .102 1.493l-.102.006h-3.197L10.037 18.5h4.213a.75.75 0 0 1 .742.648l.007.102a.75.75 0 0 1-.648.743L14.25 20h-9.5a.747.747 0 0 1-.746-.75c0-.38.28-.694.645-.743l.101-.007h3.685l.021-.065L13.45 5.499h-3.7a.75.75 0 0 1-.742-.648L9 4.75a.75.75 0 0 1 .648-.743L9.751 4h8.503h-8.503Z"/>`),
		g.Group(children),
	)
}