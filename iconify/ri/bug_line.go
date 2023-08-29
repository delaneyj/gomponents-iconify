package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 19.9a5.002 5.002 0 0 0 4-4.9v-3c0-.701-.144-1.378-.415-2h-9.17A4.981 4.981 0 0 0 7 12v3a5.002 5.002 0 0 0 4 4.9V14h2v5.9Zm-7.464-2.21A6.98 6.98 0 0 1 5 15H2v-2h3v-1c0-.643.087-1.265.249-1.856L3.036 8.866l1-1.732L6.056 8.3a7.01 7.01 0 0 1 .199-.3h11.49c.069.098.135.199.199.3l2.02-1.166l1 1.732l-2.213 1.278c.162.59.249 1.213.249 1.856v1h3v2h-3a6.96 6.96 0 0 1-.536 2.69l2.5 1.444l-1 1.732l-2.526-1.458A6.986 6.986 0 0 1 12 22a6.986 6.986 0 0 1-5.438-2.592l-2.526 1.458l-1-1.732l2.5-1.443ZM8 6a4 4 0 1 1 8 0H8Z"/>`),
		g.Group(children),
	)
}