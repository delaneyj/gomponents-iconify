package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.345 14.947H1.655c-.9 0-1.634-.703-1.634-1.57V5.616c0-.864.733-1.571 1.634-1.571h10.69c.901 0 1.635.707 1.635 1.571v7.761c-.001.867-.734 1.57-1.635 1.57zM.995 6.414v6.211l1.369 1.432h9.33l1.327-1.385V6.46l-1.327-1.524h-9.33L.995 6.414z"/><path d="M14.248 2.033H3.047l-.004.936h10.213l1.775 1.814v6.154l.896-.016V3.663c0-.87-.766-1.614-1.679-1.63zM2 9h.953v.984H2zm9 0h.984v.953H11zM5 8h.969v.969H5zm4 2h.984v.969H9z"/><path d="M8 6.016h-.969v1.015H6v.938h1.031v1.062H6v.938h1.031v1.062H6v.938h1.031v1H8v-1h.984v-.938H8V9.969h.984v-.938H8V7.969h.984v-.938H8V6.016z"/></g>`),
		g.Group(children),
	)
}