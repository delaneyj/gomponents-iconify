package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleMultipleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 4a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM2 4a3 3 0 0 1 3-3h6a3 3 0 1 1 0 6H5a3 3 0 0 1-3-3Zm3-2a2 2 0 1 0 0 4h6a2 2 0 1 0 0-4H5Zm6 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-9-1a3 3 0 0 1 3-3h6a3 3 0 1 1 0 6H5a3 3 0 0 1-3-3Zm3-2a2 2 0 1 0 0 4h6a2 2 0 1 0 0-4H5Z"/>`),
		g.Group(children),
	)
}