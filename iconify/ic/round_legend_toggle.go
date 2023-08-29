package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundLegendToggle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 15H5c-.55 0-1-.45-1-1s.45-1 1-1h14c.55 0 1 .45 1 1s-.45 1-1 1zm0 2H5c-.55 0-1 .45-1 1s.45 1 1 1h14c.55 0 1-.45 1-1s-.45-1-1-1zm-4-6l4.58-3.25c.26-.19.42-.49.42-.81c0-.81-.92-1.29-1.58-.82L15 8.55L10 5L4.48 8.36c-.3.19-.48.51-.48.86c0 .78.85 1.26 1.52.85l4.4-2.68L15 11z"/>`),
		g.Group(children),
	)
}