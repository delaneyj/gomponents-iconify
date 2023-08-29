package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brightness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.707 9.293L19 8.586V6a1 1 0 0 0-1-1h-2.586l-.707-.707l-2-2a.999.999 0 0 0-1.414 0l-2 2L8.586 5H6a1 1 0 0 0-1 1v2.586l-.707.707l-2 2a.999.999 0 0 0 0 1.414l2 2l.707.707V18a1 1 0 0 0 1 1h2.586l.707.707l2 2a.997.997 0 0 0 1.414 0l2-2l.707-.707H18a1 1 0 0 0 1-1v-2.586l.707-.707l2-2a.999.999 0 0 0 0-1.414l-2-2zm-2.414 5l-.293.293V17h-2.414l-.293.293l-1 1L12 19.586l-1.293-1.293l-1-1L9.414 17H7v-2.414l-.293-.293l-1-1L4.414 12l1.293-1.293l1-1L7 9.414V7h2.414l.293-.293l1-1L12 4.414l1.293 1.293l1 1l.293.293H17v2.414l.293.293l1 1L19.586 12l-1.293 1.293l-1 1z"/><path fill="currentColor" d="M12 8c-2.206 0-4 1.794-4 4s1.794 4 4 4s4-1.794 4-4s-1.794-4-4-4z"/>`),
		g.Group(children),
	)
}