package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideShow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7 0h1.906v2.219H7z"/><path d="M15.993 2.621c0-.334-.319-.605-.712-.605H.781c-.394 0-.712.271-.712.605v1.316h.962v8.049h4.434L3.17 14.724c-.161.162-.101.485.136.722l.43.429l3.327-3.77v3.863h1.851v-3.966l3.419 3.872l.429-.429c.237-.236.298-.56.136-.722l-2.365-2.738h4.406V3.936h1.056V2.621h-.002zm-1.962 8.485H1.959V3.937h12.072v7.169z"/></g>`),
		g.Group(children),
	)
}