package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceFiveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 32H64a32 32 0 0 0-32 32v128a32 32 0 0 0 32 32h128a32 32 0 0 0 32-32V64a32 32 0 0 0-32-32ZM92 176a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-72a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm36 36a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm36 36a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-72a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}