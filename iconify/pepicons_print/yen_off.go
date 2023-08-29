package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YenOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M11.207 9.207a1 1 0 0 1-1.414 0l-5-5a1 1 0 0 1 1.414-1.414l5 5a1 1 0 0 1 0 1.414Z"/><path d="M9.793 9.207a1 1 0 0 0 1.414 0l5-5a1 1 0 0 0-1.414-1.414l-5 5a1 1 0 0 0 0 1.414Z"/><path d="M4.5 9.5a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2h-10a1 1 0 0 1-1-1Zm0 3.5a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2h-10a1 1 0 0 1-1-1Z"/><path d="M10.5 9.5a1 1 0 0 1 1 1v7a1 1 0 1 1-2 0v-7a1 1 0 0 1 1-1Z"/></g><path fill-rule="evenodd" d="M9.336 8.87a.5.5 0 0 1-.706-.034l-5-5.512a.5.5 0 1 1 .74-.672l5 5.512a.5.5 0 0 1-.034.706Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.664 8.87a.5.5 0 0 0 .706-.034l5-5.512a.5.5 0 1 0-.74-.672l-5 5.512a.5.5 0 0 0 .034.706Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3.5 8.988a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H4a.5.5 0 0 1-.5-.5Zm0 3.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H4a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9 8a.5.5 0 0 1 .5.5v8.488a.5.5 0 0 1-1 0V8.5A.5.5 0 0 1 9 8Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}