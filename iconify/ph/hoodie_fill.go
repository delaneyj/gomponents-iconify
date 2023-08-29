package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HoodieFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m238.66 123.56l-56.3-84.44A16 16 0 0 0 169.05 32H87a16 16 0 0 0-13.31 7.12l-56.3 84.44a8 8 0 0 0-1.06 6.54l22.39 82.11A16 16 0 0 0 54.11 224H80a16 16 0 0 0 16-16v-16h64v16a16 16 0 0 0 16 16h25.89a16 16 0 0 0 15.44-11.79l22.39-82.11a8 8 0 0 0-1.06-6.54ZM80 208H54.11l-21.43-78.59L64 82.42V176a16 16 0 0 0 16 16Zm40-72a8 8 0 0 1-16 0V94.86a8 8 0 0 1 16 0Zm32-8a8 8 0 0 1-16 0V94.86a8 8 0 0 1 16 0Zm-24-49.21L83.41 53.31L87 48h82.1l3.54 5.31ZM201.89 208H176v-16a16 16 0 0 0 16-16V82.42l31.32 47Z"/>`),
		g.Group(children),
	)
}