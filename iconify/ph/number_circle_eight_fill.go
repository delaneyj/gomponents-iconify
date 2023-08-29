package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleEightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 152a20 20 0 1 1-20-20a20 20 0 0 1 20 20Zm-20-36a16 16 0 1 0-16-16a16 16 0 0 0 16 16Zm104 12A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-68 24a35.93 35.93 0 0 0-14.19-28.61a32 32 0 1 0-43.62 0A36 36 0 1 0 164 152Z"/>`),
		g.Group(children),
	)
}