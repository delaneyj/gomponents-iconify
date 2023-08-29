package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithDiagonalMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36.341" cy="36.756" r="23" fill="#fcea2b"/><circle cx="36.341" cy="36.756" r="23" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path d="M30.341 31.756a3 3 0 1 1-3-3a3.001 3.001 0 0 1 3 3m18 0a3 3 0 1 1-3-3a3.001 3.001 0 0 1 3 3"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M26 48c10-5 21-4.03 21-4.03"/>`),
		g.Group(children),
	)
}