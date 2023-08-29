package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceSixThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 36H64a28 28 0 0 0-28 28v128a28 28 0 0 0 28 28h128a28 28 0 0 0 28-28V64a28 28 0 0 0-28-28Zm20 156a20 20 0 0 1-20 20H64a20 20 0 0 1-20-20V64a20 20 0 0 1 20-20h128a20 20 0 0 1 20 20ZM100 84a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm72 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm-72 44a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm72 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm-72 44a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm72 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}