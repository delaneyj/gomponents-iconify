package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.72 2.5L2.92 6h4.33V2.5H5.72Zm3.03 0V6h4.33l-2.8-3.5H8.75Zm-6.25 11v-6h11v6h-11ZM5.48 1a1 1 0 0 0-.78.375L1.22 5.726a1 1 0 0 0-.22.625V14a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6.35a1 1 0 0 0-.22-.624l-3.48-4.35A1 1 0 0 0 10.52 1H5.48Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}