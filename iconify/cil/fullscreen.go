package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fullscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M208 48V16H16v192h32V70.627l160.687 160.686l22.626-22.626L70.627 48H208zm256 256v137.373L299.313 276.687l-22.626 22.626L441.373 464H304v32h192V304h-32z"/>`),
		g.Group(children),
	)
}