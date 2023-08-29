package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleDoubleUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204.4 51.6a108 108 0 1 0 0 152.8a108.16 108.16 0 0 0 0-152.8Zm-17 135.82a84 84 0 1 1 0-118.84a84.12 84.12 0 0 1 .02 118.84Zm-18.9-27.89a12 12 0 0 1-17 17L128 153l-23.53 23.53a12 12 0 0 1-17-17l32-32a12 12 0 0 1 17 0Zm0-56a12 12 0 1 1-17 17L128 97l-23.53 23.52a12 12 0 1 1-17-17l32-32a12 12 0 0 1 17 0Z"/>`),
		g.Group(children),
	)
}