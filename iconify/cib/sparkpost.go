package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparkpost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.599 12c-1.802 1.198-2.401 3.599-2.198 5.198c-3-3 4.599-11.396-4-17.198C20.198 7.198 8 13 8 23.198c0 4 2.599 7.604 8 8.802c5.401-1.198 8-4.802 8-8.802c0-6-3.599-8-2.401-11.198zM16 27.802c-2.401 0-4.599-2-4.599-4.599c0-2.401 2-4.604 4.599-4.604c2.401 0 4.599 2 4.599 4.604c-.198 2.599-2.198 4.599-4.599 4.599z"/>`),
		g.Group(children),
	)
}