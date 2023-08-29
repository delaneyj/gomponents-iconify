package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiSignalTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M503.785 124.254a432.019 432.019 0 0 0-495.57 0L8 124.4v27.438L86.881 264.5L237.778 480h36.444l150.9-215.5L504 151.842V124.4Zm-313.961 231.47l75.96-117.392a240.089 240.089 0 0 1 43.276 5.686l-95.22 146Zm-19.809-28.291l-38.265-54.649a238.218 238.218 0 0 1 94.9-32.873Zm63.606 90.838l107.373-164.639a239.338 239.338 0 0 1 39.256 19.152L256 450.232Zm165.018-171.748a272.034 272.034 0 0 0-285.278 0l-73.5-104.976a400.039 400.039 0 0 1 432.288 0Z"/>`),
		g.Group(children),
	)
}