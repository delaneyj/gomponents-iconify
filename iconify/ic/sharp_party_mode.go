package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPartyMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4h-5.17L15 2H9L7.17 4H2v16h20V4zM12 7c1.63 0 3.06.79 3.98 2H12c-1.66 0-3 1.34-3 3c0 .35.07.69.18 1H7.1A5.002 5.002 0 0 1 12 7zm0 10c-1.63 0-3.06-.79-3.98-2H12c1.66 0 3-1.34 3-3c0-.35-.07-.69-.18-1h2.08a5.002 5.002 0 0 1-4.9 6z"/>`),
		g.Group(children),
	)
}