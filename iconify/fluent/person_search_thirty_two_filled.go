package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSearchThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 16a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0 14c-.337 0-.67-.01-1-.027a2.492 2.492 0 0 0-.732-1.74l-2.5-2.5c.47-.98.732-2.077.732-3.233A7.466 7.466 0 0 0 13 18h13.5a3.5 3.5 0 0 1 3.5 3.5v.5c0 2.393-1.523 4.417-3.685 5.793C24.141 29.177 21.198 30 18 30ZM7 28.5c1.296 0 2.496-.41 3.476-1.11l3.317 3.317a1 1 0 0 0 1.414-1.414l-3.316-3.316A6 6 0 1 0 7 28.5Zm0-2a4 4 0 1 1 0-8a4 4 0 0 1 0 8Z"/>`),
		g.Group(children),
	)
}