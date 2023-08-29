package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jiocinema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.656 19.224H42.5v19.548a3.663 3.663 0 0 1-3.663 3.663H10.319a3.663 3.663 0 0 1-3.663-3.663V19.224h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.03 30.416l-7.759-4.868a.488.488 0 0 0-.748.414v9.736a.488.488 0 0 0 .748.414l7.759-4.868a.488.488 0 0 0 0-.828Zm7.491-24.495l-1.056 4.268l-5.225.961l1.056-4.267l-5.447 1.002l-1.057 4.268l-5.225.961l1.056-4.267l-5.447 1.002l-1.056 4.267l-5.225.962l1.056-4.267l-4.451.819a1.22 1.22 0 0 0-.98 1.422l1.136 6.173l35.252-6.487l-1.136-6.173a1.221 1.221 0 0 0-1.422-.98Z"/>`),
		g.Group(children),
	)
}