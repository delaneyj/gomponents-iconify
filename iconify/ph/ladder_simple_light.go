package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 26a6 6 0 0 0-6 6v34H70V32a6 6 0 0 0-12 0v192a6 6 0 0 0 12 0v-34h116v34a6 6 0 0 0 12 0V32a6 6 0 0 0-6-6Zm-6 52v44H70V78ZM70 178v-44h116v44Z"/>`),
		g.Group(children),
	)
}