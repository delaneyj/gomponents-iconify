package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KetoMojo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.619 27.459a3.322 3.322 0 0 1-2.285-.907L10.431 19.1a3.33 3.33 0 1 1 4.557-4.857l.012.012l5.492 5.178L33.927 5.314a3.33 3.33 0 1 1 4.824 4.59h0l-15.72 16.52a3.321 3.321 0 0 1-2.412 1.035Zm7.873.923a3.331 3.331 0 0 0-2.277-.9h-.727l-2.457 2.582a3.331 3.331 0 0 1-4.697.127l-2.873-2.71h-3.8a3.33 3.33 0 0 0-3.33 3.33v8.648a3.33 3.33 0 0 0 6.66 0V34.14h9.908l9.163 8.592a3.33 3.33 0 0 0 4.555-4.86Z"/><circle cx="20.619" cy="10.461" r="4.062" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}