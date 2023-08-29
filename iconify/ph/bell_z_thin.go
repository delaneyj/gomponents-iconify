package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellZThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 144a4 4 0 0 1-4 4h-32a4 4 0 0 1-3.33-6.22L136.53 100H112a4 4 0 0 1 0-8h32a4 4 0 0 1 3.33 6.22L119.47 140H144a4 4 0 0 1 4 4Zm70.38 46a11.84 11.84 0 0 1-10.38 6h-44.23a36 36 0 0 1-71.54 0H48a12 12 0 0 1-10.35-18C43.42 168 52 140.13 52 104a76 76 0 1 1 152 0c0 36.13 8.59 64 14.36 73.95a11.92 11.92 0 0 1 .02 12.05Zm-62.67 6h-55.42a28 28 0 0 0 55.42 0Zm55.72-14C204 169.17 196 139.31 196 104a68 68 0 1 0-136 0c0 35.32-8 65.17-15.44 78a4 4 0 0 0 0 4a3.91 3.91 0 0 0 3.44 2h160a3.91 3.91 0 0 0 3.44-2a4 4 0 0 0-.01-4Z"/>`),
		g.Group(children),
	)
}