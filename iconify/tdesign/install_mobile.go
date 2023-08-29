package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstallMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 1h11v2H6v18h12V11h2v12H4V1Zm16 0v4.657l1.53-1.466l1.384 1.445L19 9.385l-3.914-3.75l1.384-1.444L18 5.657V1h2Zm-9 16h2.004v2.004H11V17Z"/>`),
		g.Group(children),
	)
}