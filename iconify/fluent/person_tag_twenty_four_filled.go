package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonTagTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 14v2.935c0 .785.309 1.538.859 2.098l2.256 2.297c-1.18.448-2.554.671-4.115.671c-3.42 0-5.944-1.072-7.486-3.236a2.75 2.75 0 0 1-.51-1.596v-.92A2.249 2.249 0 0 1 4.253 14H11Zm4-6.995a5 5 0 1 0-10 0a5 5 0 0 0 10 0Zm1.572 15.4l-4-4.073A1.993 1.993 0 0 1 12 16.935V14a2 2 0 0 1 2.002-1.997L16.919 12a2.01 2.01 0 0 1 1.41.577l4.075 4.014a1.99 1.99 0 0 1 .012 2.829l-2.992 2.996a2.008 2.008 0 0 1-2.852-.012ZM14.996 16a1 1 0 1 0-.001-1.999a1 1 0 0 0 .001 2Z"/>`),
		g.Group(children),
	)
}