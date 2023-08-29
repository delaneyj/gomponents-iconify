package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorTextThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M180 208a4 4 0 0 1-4 4h-16a36 36 0 0 1-32-19.54A36 36 0 0 1 96 212H80a4 4 0 0 1 0-8h16a28 28 0 0 0 28-28v-44h-20a4 4 0 0 1 0-8h20V80a28 28 0 0 0-28-28H80a4 4 0 0 1 0-8h16a36 36 0 0 1 32 19.54A36 36 0 0 1 160 44h16a4 4 0 0 1 0 8h-16a28 28 0 0 0-28 28v44h20a4 4 0 0 1 0 8h-20v44a28 28 0 0 0 28 28h16a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}