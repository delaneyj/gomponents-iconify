package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coronavirus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="9.5" cy="10.5" r="1.5" fill="currentColor" opacity=".7"/><circle cx="9" cy="15" r="1" fill="currentColor"/><circle cx="14.5" cy="13.5" r="1.5" fill="currentColor" opacity=".7"/><circle cx="15" cy="9" r="1" fill="currentColor"/><path fill="currentColor" d="M12 8a1 1 0 0 1-1-1V2a1 1 0 0 1 2 0v5a1 1 0 0 1-1 1zm0 15a1 1 0 0 1-1-1v-5a1 1 0 0 1 2 0v5a1 1 0 0 1-1 1zm10-10h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2zM6 13H2a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2z" opacity=".25"/><path fill="currentColor" d="M18 13a1 1 0 0 1 0-2h2.941A9.013 9.013 0 0 0 13 3.059V7a1 1 0 0 1-2 0V3.059A9.013 9.013 0 0 0 3.059 11H6a1 1 0 0 1 0 2H3.059A9.013 9.013 0 0 0 11 20.941V17a1 1 0 0 1 2 0v3.941A9.013 9.013 0 0 0 20.941 13Zm-9 3a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm.5-4a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 9.5 12Zm5 3a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5Zm.5-5a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z" opacity=".5"/><path fill="currentColor" d="M4.979 6.393a9.063 9.063 0 0 1 1.414-1.414l-.757-.757a1 1 0 0 0-1.414 1.414zm0 11.214l-.757.757a1 1 0 1 0 1.414 1.414l.757-.757a9.063 9.063 0 0 1-1.414-1.414zm14.042 0a9.063 9.063 0 0 1-1.414 1.414l.757.757a1 1 0 0 0 1.414-1.414zm0-11.214l.757-.757a1 1 0 0 0-1.414-1.414l-.757.757a9.063 9.063 0 0 1 1.414 1.414z"/>`),
		g.Group(children),
	)
}