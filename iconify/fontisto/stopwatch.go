package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.241 6.523l1.45-1.657l-1.657-1.45l-1.45 1.553a10.446 10.446 0 0 0-4.811-1.961l-.055-.006V2.07h2.588V-.001H6.955V2.07h2.588v1.035C4.174 3.595 0 8.075 0 13.531c0 5.781 4.686 10.467 10.467 10.467s10.467-4.686 10.467-10.467c0-2.7-1.022-5.162-2.701-7.018l.008.009zm-7.662 14.806a7.972 7.972 0 1 1 0-15.944a8.131 8.131 0 0 1 5.695 2.382a7.95 7.95 0 0 1 2.381 5.683v.012v-.001c-.173 4.347-3.711 7.812-8.07 7.869h-.006z"/><path fill="currentColor" d="M9.544 7.248h2.174v7.972H9.544z"/>`),
		g.Group(children),
	)
}