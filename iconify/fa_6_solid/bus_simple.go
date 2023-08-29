package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 0c124.8 0 224 35.2 224 80v336c0 17.7-14.3 32-32 32v32c0 17.7-14.3 32-32 32h-32c-17.7 0-32-14.3-32-32v-32H128v32c0 17.7-14.3 32-32 32H64c-17.7 0-32-14.3-32-32v-32c-17.7 0-32-14.3-32-32V80C0 35.2 99.2 0 224 0zM64 128v128c0 17.7 14.3 32 32 32h256c17.7 0 32-14.3 32-32V128c0-17.7-14.3-32-32-32H96c-17.7 0-32 14.3-32 32zm16 272a32 32 0 1 0 0-64a32 32 0 1 0 0 64zm288 0a32 32 0 1 0 0-64a32 32 0 1 0 0 64z"/>`),
		g.Group(children),
	)
}