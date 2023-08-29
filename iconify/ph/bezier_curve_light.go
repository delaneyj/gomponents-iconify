package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BezierCurveLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M219.44 146.2A94.66 94.66 0 0 0 173.92 86H240a6 6 0 0 0 0-12h-82.6a30 30 0 0 0-58.8 0H16a6 6 0 0 0 0 12h66.08a94.66 94.66 0 0 0-45.52 60.2a30 30 0 1 0 12.09 1.08a82.53 82.53 0 0 1 51.4-56.39a30 30 0 0 0 55.9 0a82.53 82.53 0 0 1 51.4 56.39a30 30 0 1 0 12.09-1.08ZM58 176a18 18 0 1 1-18-18a18 18 0 0 1 18 18Zm70-78a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm88 96a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}