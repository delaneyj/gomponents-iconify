package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoleculeSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 8a3 3 0 1 0-2.707-1.705l-2.34 1.17a2.5 2.5 0 1 0 .76 3.2l2.296 1.148a2 2 0 1 0 .343-.946l-2.359-1.18a2.488 2.488 0 0 0-.338-1.456l2.223-1.11A2.99 2.99 0 0 0 11 8Zm0-1a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6.5 4a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm7.5 1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}