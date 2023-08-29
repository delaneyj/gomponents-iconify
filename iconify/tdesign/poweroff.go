package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poweroff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 2v10h-2V2h2Zm3.78 1.728l.809.589a9.5 9.5 0 1 1-11.178 0l.808-.59l1.178 1.617l-.808.59a7.5 7.5 0 1 0 8.822 0l-.808-.59l1.178-1.616Z"/>`),
		g.Group(children),
	)
}