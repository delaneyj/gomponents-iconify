package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bandlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.308 14.485l-5.034 9.33c-4.156 7.676-5.138 17.517 7.852 18.152H30.12c9.989-.414 16.42-5.939 9.469-18.152l-5.682-9.33"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.612 6.033H19.358l7.344 18.152c1.092 2.36 1.69 3.774.323 5.542c-.107.192-2.66 3.947-6.096 3.095a3.431 3.431 0 0 1-2.633-3.603a4.896 4.896 0 0 1 2.97-4.295a8.622 8.622 0 0 1 5.436-.74"/>`),
		g.Group(children),
	)
}