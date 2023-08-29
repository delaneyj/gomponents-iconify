package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarArrowRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M10.79 3.103c.495-1.004 1.926-1.004 2.421 0l2.358 4.777l5.272.766c1.108.161 1.55 1.522.749 2.303l-.906.883a6.5 6.5 0 0 0-9.442 7.43l-3.958 2.08c-.99.521-2.147-.32-1.958-1.423l.9-5.25l-3.815-3.72c-.801-.78-.359-2.142.748-2.303l5.273-.766l2.358-4.777z" fill="currentColor"/><path d="M23 17.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0zm-8.5-.5a.5.5 0 0 0 0 1h4.793l-1.647 1.646a.5.5 0 0 0 .708.708l2.5-2.5a.5.5 0 0 0 0-.708l-2.5-2.5a.5.5 0 0 0-.708.708L19.293 17H14.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}