package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 100a12 12 0 0 1 12-12h32a12 12 0 0 1 0 24h-32a12 12 0 0 1-12-12Zm136-32v128a20 20 0 0 1-20 20H40a20 20 0 0 1-20-20V68a20 20 0 0 1 20-20h36v-8a28 28 0 0 1 28-28h48a28 28 0 0 1 28 28v8h36a20 20 0 0 1 20 20ZM100 48h56v-8a4 4 0 0 0-4-4h-48a4 4 0 0 0-4 4ZM44 72v35.23A180.06 180.06 0 0 0 128 128a180 180 0 0 0 84-20.78V72Zm168 120v-58.06A204.27 204.27 0 0 1 128 152a204.21 204.21 0 0 1-84-18.06V192Z"/>`),
		g.Group(children),
	)
}