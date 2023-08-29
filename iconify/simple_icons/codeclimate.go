package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Codeclimate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.125 5.272l-4.511 4.475l2.684 2.659l1.827-1.813l5.19 5.145L24 13.079zM8.13 8.265L0 16.066l2.772 2.662l5.357-5.145l5.357 5.145l2.772-2.662z"/>`),
		g.Group(children),
	)
}