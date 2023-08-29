package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Open(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path d="M6.969 7h-1v2.031h1v-.062H7V7h-.031z"/><path fill="currentColor" d="M5.969 7h1v-.979L5.031 6v5.958h.938v-2.02h1v-.907h-1V7zM7 7h1v1.969H7z"/><path d="M2.984 7.016v-.021H1v.021H.984v3.937H1v.078h1.984v-.078H3V7.016h-.016z"/><path fill="currentColor" d="M3 7h.947v3.938H3zM1 6h1.984v.943H1zm0 5h1.984v.953H1zM0 7h.949v3.938H0zm15.031-.969V7.5l-1.062-.875v-.594h-.953v5.938h.953V8.031l1.062.875v3.063h.922V6.031h-.922zm-4.947-.005H9.016v5.963h1.068v-.02h1.885v-.938H9.953V8.969h2.016v-.938H9.953V6.953h2.016v-.906h-1.885v-.021z"/></g>`),
		g.Group(children),
	)
}