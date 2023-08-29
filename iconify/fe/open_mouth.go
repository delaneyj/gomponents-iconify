package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feOpenMouth0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feOpenMouth1" fill="currentColor" fill-rule="nonzero"><path id="feOpenMouth2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm2.5-6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/></g></g>`),
		g.Group(children),
	)
}