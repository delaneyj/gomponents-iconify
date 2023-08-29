package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceSixBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 28H64a36 36 0 0 0-36 36v128a36 36 0 0 0 36 36h128a36 36 0 0 0 36-36V64a36 36 0 0 0-36-36Zm12 164a12 12 0 0 1-12 12H64a12 12 0 0 1-12-12V64a12 12 0 0 1 12-12h128a12 12 0 0 1 12 12ZM112 84a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm64 0a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm-64 44a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm64 0a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm-64 44a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm64 0a16 16 0 1 1-16-16a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}