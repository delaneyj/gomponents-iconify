package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoOffFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m31 32.768l10.866 10.866a1.25 1.25 0 0 0 1.768-1.768l-37.5-37.5a1.25 1.25 0 1 0-1.768 1.768l3.386 3.385A6.251 6.251 0 0 0 4 15.25v17.5A6.25 6.25 0 0 0 10.25 39h14.5A6.25 6.25 0 0 0 31 32.768Zm2-5.07l7.66 7.66c1.547.56 3.34-.559 3.34-2.36V15.003c0-2.083-2.397-3.252-4.039-1.97L33 18.468v9.228ZM14.303 9L31 25.697V15.25A6.25 6.25 0 0 0 24.75 9H14.303Z"/>`),
		g.Group(children),
	)
}