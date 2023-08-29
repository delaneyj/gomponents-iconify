package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserRectangleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 36H40a20 20 0 0 0-20 20v144a20 20 0 0 0 20 20h176a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm-116 92a28 28 0 1 1 28 28a28 28 0 0 1-28-28Zm28 52a59.66 59.66 0 0 1 40.85 16h-81.7A59.66 59.66 0 0 1 128 180Zm84 16h-12.44A83.46 83.46 0 0 0 165 164.5a52 52 0 1 0-74 0A83.46 83.46 0 0 0 56.44 196H44V60h168Z"/>`),
		g.Group(children),
	)
}