package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogInFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaLogInFill0"><g id="evaLogInFill1"><path id="evaLogInFill2" fill="currentColor" d="M19 4h-2a1 1 0 0 0 0 2h1v12h-1a1 1 0 0 0 0 2h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Zm-7.2 3.4a1 1 0 0 0-1.6 1.2L12 11H4a1 1 0 0 0 0 2h8.09l-1.72 2.44a1 1 0 0 0 .24 1.4a1 1 0 0 0 .58.18a1 1 0 0 0 .81-.42l2.82-4a1 1 0 0 0 0-1.18Z"/></g></g>`),
		g.Group(children),
	)
}