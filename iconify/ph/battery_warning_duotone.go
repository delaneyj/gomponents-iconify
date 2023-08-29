package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryWarningDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 80v96a16 16 0 0 1-16 16H32a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h168a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M200 56H32A24 24 0 0 0 8 80v96a24 24 0 0 0 24 24h168a24 24 0 0 0 24-24V80a24 24 0 0 0-24-24Zm8 120a8 8 0 0 1-8 8H32a8 8 0 0 1-8-8V80a8 8 0 0 1 8-8h168a8 8 0 0 1 8 8Zm48-80v64a8 8 0 0 1-16 0V96a8 8 0 0 1 16 0Zm-148 28V96a8 8 0 0 1 16 0v28a8 8 0 0 1-16 0Zm20 32a12 12 0 1 1-12-12a12 12 0 0 1 12 12Z"/></g>`),
		g.Group(children),
	)
}