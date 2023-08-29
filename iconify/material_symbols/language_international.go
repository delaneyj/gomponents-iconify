package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageInternational(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 16V8h2v8H1Zm3 0V8h2l2.225 4.45L8 12.5V8h2v8H8l-2.225-4.45L6 11.5V16H4Zm9 0v-6h-2V8h6v2h-2v6h-2Zm5 0V8h2v6h3v2h-5Z"/>`),
		g.Group(children),
	)
}