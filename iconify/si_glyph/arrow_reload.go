package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowReload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.796 8.908L15.234 7.21a.553.553 0 0 0-.776 0l-1.564 1.698a.543.543 0 0 0 0 .772h1.294a5.345 5.345 0 0 1-3.789 3.79a5.378 5.378 0 0 1-5.767-2.119l-1.091.751a6.709 6.709 0 0 0 7.196 2.643A6.665 6.665 0 0 0 15.55 9.68h1.245a.544.544 0 0 0 .001-.772zM5.475 8.021a.543.543 0 0 0 0-.772H4.018a5.339 5.339 0 0 1 3.771-3.738a5.373 5.373 0 0 1 5.766 2.121l1.092-.752a6.712 6.712 0 0 0-7.199-2.645a6.67 6.67 0 0 0-4.8 5.014H1.196a.543.543 0 0 0 0 .772l1.638 1.637a.553.553 0 0 0 .776 0l1.865-1.637z"/>`),
		g.Group(children),
	)
}