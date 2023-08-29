package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.016 3.984h4.981L10.016.016v3.968z"/><path d="M8.969 5.016V.011H3.034v15.958h11.942V5.016H8.969zm3.058 6.642c-.045.379-.566.935-1.07 1.143c-.656.271-1.34.12-1.528-.336c-.188-.456.19-1.046.847-1.316c.231-.096.48-.15.709-.158V8.363l-2.973.699l.004 3.548c-.045.379-.649 1.02-1.153 1.228c-.655.271-1.339.12-1.529-.336c-.188-.456.192-1.046.848-1.316c.28-.116.528-.19.788-.166V8.142l5.062-1.188l-.005 4.704z"/></g>`),
		g.Group(children),
	)
}