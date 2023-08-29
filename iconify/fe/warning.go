package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feWarning0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWarning1" fill="currentColor" fill-rule="nonzero"><path id="feWarning2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm-1-6h2v2h-2v-2Zm0-10h2v8h-2V6Z"/></g></g>`),
		g.Group(children),
	)
}