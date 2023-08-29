package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePrint0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePrint1" fill="currentColor"><path id="fePrint2" d="M7 11V4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v7h1a2 2 0 0 1 2 2v3c0 1.1-.9 2-2 2h-1v2a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-2H6c-1.1 0-2-.9-2-2v-3a2 2 0 0 1 2-2h1Zm2-7v7h6V4H9Zm5 9v2h2v-2h-2Zm-5 5v2h6v-2H9Z"/></g></g>`),
		g.Group(children),
	)
}