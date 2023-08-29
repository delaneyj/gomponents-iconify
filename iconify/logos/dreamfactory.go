package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dreamfactory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M102.602 165.683C77.852 133.2 92.582 67.731 155.417 72.5c-74.78-86.207-171.265 63.165-52.815 93.183" fill="#B83C23"/><path d="M213.277.175L256 25.74V237l-2.427-.58l-39.283-22.377l-1.013-1.6V.175z" fill="#F78D2A"/><path d="M256 237H43.396L0 212.443h213.277L256 237z" fill="#B83C23"/>`),
		g.Group(children),
	)
}