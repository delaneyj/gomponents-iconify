package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryWarningVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 136V96a8 8 0 0 1 16 0v40a8 8 0 0 1-16 0Zm8 24a12 12 0 1 0 12 12a12 12 0 0 0-12-12ZM96 16h64a8 8 0 0 0 0-16H96a8 8 0 0 0 0 16Zm104 40v168a24 24 0 0 1-24 24H80a24 24 0 0 1-24-24V56a24 24 0 0 1 24-24h96a24 24 0 0 1 24 24Zm-16 0a8 8 0 0 0-8-8H80a8 8 0 0 0-8 8v168a8 8 0 0 0 8 8h96a8 8 0 0 0 8-8Z"/>`),
		g.Group(children),
	)
}