package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 44h-60V16a12 12 0 0 0-24 0v28H56a36 36 0 0 0-36 36v112a36 36 0 0 0 36 36h144a36 36 0 0 0 36-36V80a36 36 0 0 0-36-36Zm12 148a12 12 0 0 1-12 12H56a12 12 0 0 1-12-12V80a12 12 0 0 1 12-12h144a12 12 0 0 1 12 12Zm-48-64H92a32 32 0 0 0 0 64h72a32 32 0 0 0 0-64Zm-28 24v16h-16v-16Zm-52 8a8 8 0 0 1 8-8h4v16h-4a8 8 0 0 1-8-8Zm80 8h-4v-16h4a8 8 0 0 1 0 16Zm-96-68a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm88 0a16 16 0 1 1 16 16a16 16 0 0 1-16-16Z"/>`),
		g.Group(children),
	)
}