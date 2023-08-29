package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8.226 7.232a1 1 0 0 1 1.409.128l2.083 2.5a1 1 0 1 1-1.536 1.28l-2.084-2.5a1 1 0 0 1 .128-1.408Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.226 13.768a1 1 0 0 1-.128-1.408l2.084-2.5a1 1 0 1 1 1.536 1.28l-2.083 2.5a1 1 0 0 1-1.409.128Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.7 10.5a1 1 0 0 1-1 1H3.2a1 1 0 1 1 0-2h6.5a1 1 0 0 1 1 1Zm5.55 7a1 1 0 0 1-1 1h-9a1 1 0 1 1 0-2h9a1 1 0 0 1 1 1Zm0-14a1 1 0 0 1-1 1h-9a1 1 0 1 1 0-2h9a1 1 0 0 1 1 1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.25 18.5a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1Zm0-10a1 1 0 0 1-1-1v-4a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1Zm9.35 10a1 1 0 0 1-1-1v-14a1 1 0 1 1 2 0v14a1 1 0 0 1-1 1Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}