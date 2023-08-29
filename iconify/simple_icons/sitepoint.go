package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sitepoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.471 10.533l1.771 1.688l5.598 5.141l2.4-2.291a.804.804 0 0 0-.046-.985L9.99 12.184l.01-.005l-2.371-2.266a.788.788 0 0 1 .021-1.079l6.39-6.076L11.146 0L2.475 8.238a1.571 1.571 0 0 0 0 2.295h-.004zm19.056 2.937l-1.77-1.691l-5.595-5.142l-2.411 2.291a.773.773 0 0 0 .045.985l2.205 1.891h-.006l2.369 2.265a.77.77 0 0 1-.029 1.064l-6.391 6.075L12.855 24l8.67-8.238a1.57 1.57 0 0 0 0-2.295l.002.003z"/>`),
		g.Group(children),
	)
}