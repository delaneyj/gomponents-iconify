package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.061 11.077l-17.32-6.92a.994.994 0 0 0-1.17.32a1 1 0 0 0-.01 1.22l4.49 6a.525.525 0 0 1-.01.62L2.511 18.3a1.02 1.02 0 0 0 0 1.22a1 1 0 0 0 .8.4a1.021 1.021 0 0 0 .38-.07l17.36-6.9a1.006 1.006 0 0 0 .01-1.87Zm-17.69-5.99l16.06 6.42H8.061a1.329 1.329 0 0 0-.21-.41Zm-.06 13.82l4.53-5.98a1.212 1.212 0 0 0 .22-.42h11.38Z"/>`),
		g.Group(children),
	)
}