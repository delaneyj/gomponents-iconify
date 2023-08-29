package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideRecordSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2.6a5.477 5.477 0 0 1-.578-3H4.5a.5.5 0 0 1 0-1h.707c.099-.349.23-.683.393-1H4.5a.5.5 0 0 1 0-1h1.757a5.54 5.54 0 0 1 1.08-1H4.5a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .489.605A5.477 5.477 0 0 1 10.5 5c1.86 0 3.505.923 4.5 2.337V5a2 2 0 0 0-2-2H3Zm7.5 11a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0 1a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm0-2a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/>`),
		g.Group(children),
	)
}