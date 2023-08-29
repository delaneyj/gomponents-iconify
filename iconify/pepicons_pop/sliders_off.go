package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2.5 4.75a1 1 0 0 1 1-1H10a1 1 0 1 1 0 2H3.5a1 1 0 0 1-1-1Zm11.375 0a1 1 0 0 1 1-1H16.5a1 1 0 1 1 0 2h-1.625a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.75 5.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Zm-10.25 7a1 1 0 0 1 1-1H10a1 1 0 1 1 0 2H3.5a1 1 0 0 1-1-1Zm11.375 0a1 1 0 0 1 1-1H16.5a1 1 0 1 1 0 2h-1.625a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.75 15.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Zm-10.25-8a1 1 0 0 1 1-1h1.625a1 1 0 0 1 0 2H3.5a1 1 0 0 1-1-1Zm6.5 0a1 1 0 0 1 1-1h6.5a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.75 10.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}