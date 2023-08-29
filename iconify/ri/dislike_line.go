package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DislikeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.808 1.394l18.384 18.384l-1.414 1.415l-3.746-3.747L12 21.486l-8.478-8.493a6 6 0 0 1 .033-8.023L1.394 2.808l1.414-1.414Zm2.172 10.23l7.02 7.03l2.618-2.622l-9.646-9.646a4 4 0 0 0 .008 5.237Zm15.263-6.866a6 6 0 0 1 .236 8.235l-1.635 1.636l-1.414-1.414l1.59-1.592a4 4 0 0 0-5.683-5.606l-1.335 1.198l-1.336-1.197a3.975 3.975 0 0 0-1.155-.723l-2.25-2.25A5.99 5.99 0 0 1 12 4.53a5.998 5.998 0 0 1 8.242.229Z"/>`),
		g.Group(children),
	)
}