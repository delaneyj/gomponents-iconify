package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileNew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M255.583 0L29.513 218.249V1200H635.49v-104.328H134.506V292.179h197.589v-187.85h442.844v312.986h104.992V0H255.583zm472.725 508.73v249.091H479.145V950.91h249.163V1200h193.016V950.91h249.164V757.821H921.323V508.73H728.308z"/>`),
		g.Group(children),
	)
}