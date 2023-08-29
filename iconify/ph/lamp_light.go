package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m245.51 149.64l-48-112A6 6 0 0 0 192 34H64a6 6 0 0 0-5.51 3.64l-48 112A6 6 0 0 0 16 158h106v52H96a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12h-26v-52h60v34a6 6 0 0 0 12 0v-34h34a6 6 0 0 0 5.51-8.36ZM25.1 146L68 46h120l42.9 100Z"/>`),
		g.Group(children),
	)
}