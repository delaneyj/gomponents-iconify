package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyEthThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m219.15 125.53l-88-112a4 4 0 0 0-6.3 0l-88 112a4 4 0 0 0 0 4.94l88 112a4 4 0 0 0 6.3 0l88-112a4 4 0 0 0 0-4.94ZM132 27.57l77.71 98.9L132 161.79Zm-8 134.22l-77.71-35.32L124 27.57Zm0 8.79v57.85l-70.72-90Zm8 0l70.72-32.15l-70.72 90Z"/>`),
		g.Group(children),
	)
}