package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitBranchBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 64a36 36 0 1 0-48 33.94V104a12 12 0 0 1-12 12H96a35.88 35.88 0 0 0-12 2.06V97.94a36 36 0 1 0-24 0v60.12a36 36 0 1 0 24 0V152a12 12 0 0 1 12-12h72a36 36 0 0 0 36-36v-6.06A36.07 36.07 0 0 0 228 64ZM72 52a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm0 152a12 12 0 1 1 12-12a12 12 0 0 1-12 12ZM192 76a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}