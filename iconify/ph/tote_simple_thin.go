package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToteSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M233 72.06a12.11 12.11 0 0 0-9.08-4.06H172v-4a44 44 0 0 0-88 0v4H32.08A12.11 12.11 0 0 0 23 72.06a12 12 0 0 0-2.94 9.35l14.26 120a12 12 0 0 0 12 10.59h163.35a12 12 0 0 0 12-10.59l14.26-120a12 12 0 0 0-2.93-9.35ZM92 64a36 36 0 0 1 72 0v4H92Zm136 16.46l-14.25 120a4 4 0 0 1-4.05 3.54H46.33a4 4 0 0 1-4-3.54L28 80.46a3.9 3.9 0 0 1 1-3.09A4.08 4.08 0 0 1 32.08 76h191.84a4.08 4.08 0 0 1 3.08 1.37a3.9 3.9 0 0 1 1 3.09Z"/>`),
		g.Group(children),
	)
}