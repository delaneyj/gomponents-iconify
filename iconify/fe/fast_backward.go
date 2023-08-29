package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFastBackward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFastBackward1" fill="currentColor"><path id="feFastBackward2" d="M5 12c0-.107.023-.216.072-.316a.617.617 0 0 1 .221-.253l6.86-4.35c.276-.174.623-.06.775.254a.725.725 0 0 1 .072.316V12c0-.106.023-.215.072-.315a.617.617 0 0 1 .221-.253l6.86-4.35c.276-.174.623-.06.775.254a.725.725 0 0 1 .072.316v8.698c0 .36-.255.651-.57.651a.516.516 0 0 1-.277-.082l-6.86-4.349A.671.671 0 0 1 13 12v4.35c0 .36-.255.651-.57.651a.516.516 0 0 1-.277-.082l-6.86-4.349A.671.671 0 0 1 5 12v4c0 .552-.5 1-1 1s-1-.448-1-1V8a1 1 0 1 1 2 0v4Z"/></g></g>`),
		g.Group(children),
	)
}