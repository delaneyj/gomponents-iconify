package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UniversalAccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4C9.383 4 4 9.383 4 16s5.383 12 12 12s12-5.383 12-12S22.617 4 16 4zm0 2c5.535 0 10 4.465 10 10s-4.465 10-10 10S6 21.535 6 16S10.465 6 16 6zm0 2a1.999 1.999 0 1 0 0 4a1.999 1.999 0 1 0 0-4zm-5.781 4.031L9.78 13.97l4.219.937V19l-.969 4.813l1.938.375l1-5l.031-.094l.031.093l1 5l1.938-.375L18 19v-4.094l4.219-.937l-.438-1.938L17.5 13h-3z"/>`),
		g.Group(children),
	)
}