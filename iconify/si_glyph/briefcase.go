package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M6.959 3.079c0-.619.482-1.124 1.072-1.124h1.838c.59 0 1.184.505 1.184 1.124v-.058h.919v-.474c0-.86-.478-1.548-1.458-1.548H7.385c-.654 0-1.362.73-1.362 1.548v.474h.937v.058h-.001z"/><path d="M11 5.969v1.012h6V3.016H.985v3.965h6.003V5.969h4.013zM7 8.998v-.995H1.022V16H17V8.003h-5.987v.995H7.001zM8 7v.967h1.99V7H8z"/></g>`),
		g.Group(children),
	)
}