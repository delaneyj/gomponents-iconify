package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.846 8.886L24 2.932v13.82L9.621 21.068L0 14.09l3.35-9.956l7.496 4.751v.001zm-4.582 2.067l3.923-1.768l-6.065-3.85l2.142 5.618zm-5.393 2.44l4.842-2.187l-2.179-5.717l-2.662 7.904H.871zm22.526 2.54V4.256l-5.96 7.37l5.96 4.307zm-12.865 4.233l12.497-3.758l-5.973-4.316l-6.524 8.074zM.94 14.029l8.092 5.867l-3.106-8.124L.94 14.029zm21.722-9.826c-5.085 2.296-10.163 4.6-15.25 6.895l9.445.284l5.805-7.178v-.001zM9.775 20.143l6.608-8.173l-9.844-.29l3.236 8.462v.001z"/>`),
		g.Group(children),
	)
}