package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 24c3-5.455 11.2-13.745 20-9.818C33 9.818 41 18.545 44 24c-3 4.364-11.2 12-20 12S7 28.364 4 24Zm5.31-.477C12.414 22.473 18.183 21 24 21c5.747 0 11.45 1.439 14.58 2.486c.457.152.431.798-.03.936l-.025.008C35.5 25.34 29.98 27 24 27c-6.08 0-11.683-1.625-14.672-2.573c-.448-.142-.463-.754-.019-.904Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}