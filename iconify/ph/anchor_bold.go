package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 132a12 12 0 0 0-12 12a36 36 0 0 1-36 36a51.7 51.7 0 0 0-28 8.22V132h28a12 12 0 0 0 0-24h-28V89.94a36 36 0 1 0-24 0V108H88a12 12 0 0 0 0 24h28v56.22A51.7 51.7 0 0 0 88 180a36 36 0 0 1-36-36a12 12 0 0 0-24 0a60.07 60.07 0 0 0 60 60a28 28 0 0 1 28 28a12 12 0 0 0 24 0a28 28 0 0 1 28-28a60.07 60.07 0 0 0 60-60a12 12 0 0 0-12-12Zm-88-88a12 12 0 1 1-12 12a12 12 0 0 1 12-12Z"/>`),
		g.Group(children),
	)
}