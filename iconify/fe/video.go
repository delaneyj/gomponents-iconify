package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feVideo0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feVideo1" fill="currentColor"><path id="feVideo2" d="m15 9.649l5.646-2.137A1 1 0 0 1 22 8.448v7.109a1 1 0 0 1-1.351.936L15 14.375V16a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1.649Z"/></g></g>`),
		g.Group(children),
	)
}