package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitBranchThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 64a28 28 0 1 0-32 27.71V104a20 20 0 0 1-20 20H96a27.9 27.9 0 0 0-20 8.43V91.71a28 28 0 1 0-8 0v72.58a28 28 0 1 0 8 0V152a20 20 0 0 1 20-20h72a28 28 0 0 0 28-28V91.71A28 28 0 0 0 220 64ZM52 64a20 20 0 1 1 20 20a20 20 0 0 1-20-20Zm40 128a20 20 0 1 1-20-20a20 20 0 0 1 20 20ZM192 84a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}