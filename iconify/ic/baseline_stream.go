package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineStream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="20" cy="12" r="2" fill="currentColor"/><circle cx="4" cy="12" r="2" fill="currentColor"/><circle cx="12" cy="20" r="2" fill="currentColor"/><path fill="currentColor" d="M10.05 8.59L6.03 4.55h-.01l-.31-.32l-1.42 1.41l4.02 4.05l.01-.01l.31.32zm3.893.027l4.405-4.392L19.76 5.64l-4.405 4.393zM10.01 15.36l-1.42-1.41l-4.03 4.01l-.32.33l1.41 1.41l4.03-4.02zm9.75 2.94l-3.99-4.01l-.36-.35L14 15.35l3.99 4.01l.35.35z"/><circle cx="12" cy="4" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}