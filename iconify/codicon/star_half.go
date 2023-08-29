package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.405 6.252L8 1l1.595 5.252H15l-4.373 3.4L12.25 15L8 11.695L3.75 15l1.623-5.348L1 6.252h5.405zM8 10.032l1.915 1.49l-.731-2.41l1.915-1.49H8.732L8 5.214v4.82zm0-7.525zm5.652 4.215H9.28h4.372z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}