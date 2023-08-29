package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBriefcaseFill0"><g id="evaBriefcaseFill1"><path id="evaBriefcaseFill2" fill="currentColor" d="M7 21h10V7h-1V5.5A2.5 2.5 0 0 0 13.5 3h-3A2.5 2.5 0 0 0 8 5.5V7H7Zm3-15.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5V7h-4ZM19 7v14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3ZM5 7a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3Z"/></g></g>`),
		g.Group(children),
	)
}