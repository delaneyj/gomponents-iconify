package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mutify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.577 16.894l8.845.014m-8.845-.014l-.027 11.6l8.898.04l-.026-11.626"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.55 28.493l8.814 13.62l3.245.028l-3.16-13.607l14.97 11.36L29.375 5.5L14.422 16.908m17.131.145a4.21 4.21 0 0 1 3.777 2.767a7.131 7.131 0 0 1 .004 5.66a4.293 4.293 0 0 1-3.773 2.831"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.7 12.234c2.572-.037 4.961 1.947 6.256 5.197a14.738 14.738 0 0 1 .007 10.517c-1.292 3.252-3.678 5.24-6.25 5.205"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.652 8.259c3.844-.046 7.412 2.722 9.346 7.25a19.214 19.214 0 0 1 .004 14.642c-1.931 4.524-5.498 7.286-9.342 7.233M10.326 6.16L38.028 42.5"/>`),
		g.Group(children),
	)
}