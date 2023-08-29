package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M211.89 119.09a4 4 0 0 0-2.49-2.84l-60.81-22.8l15.33-76.67a4 4 0 0 0-6.84-3.51l-112 120a4 4 0 0 0-1 3.64a4 4 0 0 0 2.49 2.84l60.81 22.8l-15.3 76.67a4 4 0 0 0 6.84 3.51l112-120a4 4 0 0 0 .97-3.64ZM102.68 227l13.24-66.2a4 4 0 0 0-2.52-4.53L55 134.36L153.32 29l-13.24 66.2a4 4 0 0 0 2.52 4.53l58.4 21.91Z"/>`),
		g.Group(children),
	)
}