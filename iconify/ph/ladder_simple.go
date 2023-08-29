package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 24a8 8 0 0 0-8 8v32H72V32a8 8 0 0 0-16 0v192a8 8 0 0 0 16 0v-32h112v32a8 8 0 0 0 16 0V32a8 8 0 0 0-8-8Zm-8 56v40H72V80ZM72 176v-40h112v40Z"/>`),
		g.Group(children),
	)
}