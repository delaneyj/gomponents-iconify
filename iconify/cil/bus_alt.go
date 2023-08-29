package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M441.884 109.647A32.029 32.029 0 0 0 415.669 96H48a32.036 32.036 0 0 0-32 32v248h41.082a67.982 67.982 0 0 0 133.836 0h122.164a67.982 67.982 0 0 0 133.836 0H496V194.521a23.886 23.886 0 0 0-4.338-13.763ZM180 128h144v96H180Zm-132 0h100v96H48Zm76 272a36 36 0 1 1 36-36a36.04 36.04 0 0 1-36 36Zm256 0a36 36 0 1 1 36-36a36.04 36.04 0 0 1-36 36Zm84-56h-19.006a68 68 0 0 0-129.988 0H188.994a68 68 0 0 0-129.988 0H48v-88h416Zm0-120H356v-96h59.669L464 197.043Z"/>`),
		g.Group(children),
	)
}