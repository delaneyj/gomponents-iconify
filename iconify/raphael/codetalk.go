package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Codetalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4.938c-7.732 0-14 4.7-14 10.5c0 1.98.74 3.833 2.016 5.414L2 25.272l5.613-1.44c2.34 1.316 5.237 2.106 8.387 2.106c7.732 0 14-4.7 14-10.5s-6.268-10.5-14-10.5zM13.704 19.47l-2.338 2.336l-6.43-6.43l6.43-6.433l2.338 2.34l-4.09 4.092l4.09 4.095zm7.07 2.333l-2.336-2.34l4.092-4.09l-4.092-4.09l2.337-2.34l6.43 6.426l-6.43 6.433z"/>`),
		g.Group(children),
	)
}