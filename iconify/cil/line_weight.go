package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineWeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 464h480v32H16zm0-32h480v-80H16Zm32-48h416v16H48Zm-32-64h480V208H16Zm32-80h416v48H48Zm-32-64h480V16H16ZM48 48h416v96H48Z"/>`),
		g.Group(children),
	)
}