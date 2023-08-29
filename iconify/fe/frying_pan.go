package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FryingPan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFryingPan0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFryingPan1" fill="currentColor" fill-rule="nonzero"><path id="feFryingPan2" d="M13 18.268a2 2 0 1 1-2 0v-2.339A7.002 7.002 0 0 1 12 2a7 7 0 0 1 1 13.93v2.338ZM12 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></g>`),
		g.Group(children),
	)
}