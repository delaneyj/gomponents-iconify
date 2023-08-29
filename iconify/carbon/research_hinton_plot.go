package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResearchHintonPlot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 2h4v4H2zm8 0h4v4h-4zm8 0h4v4h-4zm8 0h4v4h-4zM2 10h4v4H2zm8 0h4v4h-4zm8 0h4v4h-4zm8 0h4v4h-4zM2 18h4v4H2zm8 0h4v4h-4zm8 0h4v4h-4zm8 0h4v4h-4zM2 26h4v4H2zm8 0h4v4h-4zm8 0h4v4h-4zm8 0h4v4h-4z"/>`),
		g.Group(children),
	)
}