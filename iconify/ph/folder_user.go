package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M214.61 198.62a32 32 0 1 0-45.23 0a40 40 0 0 0-17.11 23.32a8 8 0 0 0 5.67 9.79a8.15 8.15 0 0 0 2.06.27a8 8 0 0 0 7.73-5.95C170.56 215.42 180.54 208 192 208s21.44 7.42 24.27 18.05a8 8 0 1 0 15.46-4.11a40 40 0 0 0-17.12-23.32ZM192 160a16 16 0 1 1-16 16a16 16 0 0 1 16-16Zm24-88h-84.69L104 44.69A15.86 15.86 0 0 0 92.69 40H40a16 16 0 0 0-16 16v144.61A15.4 15.4 0 0 0 39.38 216h81.18a8 8 0 0 0 0-16H40V88h176v32a8 8 0 0 0 16 0V88a16 16 0 0 0-16-16ZM92.69 56l16 16H40V56Z"/>`),
		g.Group(children),
	)
}