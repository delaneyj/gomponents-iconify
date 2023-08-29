package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderContact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7 8v1.957h2.014V8H7zm3-1h1v3h-1z"/><path d="M7.35 4L5.788 2.042H2.021v1.021H.023v9.913h1.02l.002 1h14.902L15.968 4H7.35zm3.643 6l-.018 1L7 11.067v-.99l-1-.021V7.972h1v-.98l3-.01V6H6v1h-.984v3.964h1v1.007h3.957V13H6v-1l-1 .058v-1.021H3.953V7.034h1.008V6.025H6V5h3.958v1.011H11v.981h.953V10h-.96z"/><path d="M12.958 2.979h-4.27l-.667-.958h4.937v.958z"/></g>`),
		g.Group(children),
	)
}