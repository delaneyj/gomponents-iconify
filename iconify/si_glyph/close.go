package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Close(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.053 10.938v-.906h.895v.906h-.895zM1.048 9.932V7.037h.905v2.895h-.905zm1.005-2.994v-.906h.895v.906h-.895zm5.968 4.041V10H9v.979h-.979zm1.016-1.038V7.037h.905v2.904h-.905zm-1.978 0V7.037h.873v2.904h-.873zM8 6.973v-.972h.951v.972H8zm-3.941 3.958V6.059h.883v3.989h1.031v.884l-1.914-.001zm10.999.016V6.042h1.9v.905h-1.04V8.1h1.061v.847H15.94v1.106h1.018v.894h-1.9zm-3.979-.031v-.832h1.863v.832h-1.863zm1.928-.969v-.921h.936v.921h-.936zm-.976-.983v-.943h.932v.943h-.932zm-.994-1.006v-.932h.926v.932h-.926zm1.042-1.042v-.848h1.863v.848h-1.863z"/>`),
		g.Group(children),
	)
}