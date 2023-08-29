package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorParagraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 2H7.54c-.83 0-1.59.2-2.28.6c-.7.41-1.25.96-1.65 1.65C3.2 4.94 3 5.7 3 6.52s.2 1.58.61 2.27c.4.69.95 1.24 1.65 1.64c.69.41 1.45.61 2.28.61h.43V17c0 .27.1.51.29.71c.2.19.44.29.71.29c.28 0 .51-.1.71-.29c.2-.2.3-.44.3-.71V5c0-.27.09-.51.29-.71c.2-.19.44-.29.71-.29s.51.1.71.29c.19.2.29.44.29.71v12c0 .27.1.51.3.71c.2.19.43.29.71.29c.27 0 .51-.1.71-.29c.19-.2.29-.44.29-.71V4H15c.27 0 .5-.1.7-.3c.2-.19.3-.43.3-.7s-.1-.51-.3-.71A.984.984 0 0 0 15 2z"/>`),
		g.Group(children),
	)
}