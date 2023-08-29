package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 9.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM7.5 13a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 10.75a.75.75 0 0 1 .75.75V13a.75.75 0 0 1-1.5 0v-1.5a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.25 13a.75.75 0 0 1-.75.75H13a.75.75 0 0 1 0-1.5h1.5a.75.75 0 0 1 .75.75ZM10.5 5.5a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11.775 4.538a1 1 0 0 1 .687 1.237l-1 3.5a1 1 0 1 1-1.924-.55l1-3.5a1 1 0 0 1 1.237-.687Zm0 16.924a1 1 0 0 0 .687-1.237l-1-3.5a1 1 0 1 0-1.924.55l1 3.5a1 1 0 0 0 1.237.687Zm2.45 0a1 1 0 0 1-.686-1.237l1-3.5a1 1 0 1 1 1.923.55l-1 3.5a1 1 0 0 1-1.237.687Zm0-16.924a1 1 0 0 0-.686 1.237l1 3.5a1 1 0 1 0 1.923-.55l-1-3.5a1 1 0 0 0-1.237-.687Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.5 20.5a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}