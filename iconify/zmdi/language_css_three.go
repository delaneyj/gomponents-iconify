package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageCssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M59 30h342l-31 156l-5 25l-24 121l-183 61L0 332l16-80h67l-6 33l95 36l111-36l15-77H24l13-67h274l9-44H46z"/>`),
		g.Group(children),
	)
}