package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinAssistant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 3q18 0 30.5 12.5T384 45v299q0 18-12.5 30.5T341 387h-85l-64 64l-64-64H43q-18 0-30.5-12.5T0 344V45q0-17 12.5-29.5T43 3h298zM232 235l88-40l-88-40l-40-88l-40 88l-88 40l88 40l40 88z"/>`),
		g.Group(children),
	)
}