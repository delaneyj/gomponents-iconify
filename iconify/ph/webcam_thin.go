package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebcamThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 104a36 36 0 1 0-36 36a36 36 0 0 0 36-36Zm-64 0a28 28 0 1 1 28 28a28 28 0 0 1-28-28Zm124 100h-92v-24.11a76 76 0 1 0-8 0V204H32a4 4 0 0 0 0 8h192a4 4 0 0 0 0-8ZM60 104a68 68 0 1 1 68 68a68.07 68.07 0 0 1-68-68Z"/>`),
		g.Group(children),
	)
}