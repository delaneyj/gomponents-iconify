package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 12h4v4H1v-4zm5 0h4v4H6v-4zm5 0h4v4h-4v-4zM1 7h4v4H1V7zm0-5h4v4H1V2zm5 5h4v4H6V7zm1-6h4v4H7V1zm4 6h4v4h-4V7zm2-7h3v3h-3V0z"/>`),
		g.Group(children),
	)
}