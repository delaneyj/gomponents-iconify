package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0L0 256l42.7 42.7l64-64V512h298.7V234.7l64 64L512 256L256 0zm-21.3 448h-64v-64h64v64zm0-106.7h-64v-64h64v64zM341.3 448h-64v-64h64v64zm0-106.7h-64v-64h64v64zM426.7 0h-64v42.7l64 64V0z"/>`),
		g.Group(children),
	)
}