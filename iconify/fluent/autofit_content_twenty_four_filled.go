package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutofitContentTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M6.064 5.244a.75.75 0 1 0-1.128-.988l-1.75 2a.75.75 0 0 0 0 .988l1.75 2a.75.75 0 0 0 1.128-.988L5.403 7.5h2.862a.75.75 0 0 0 0-1.5H5.403l.661-.756z" fill="currentColor"/><path d="M18.02 4.186a.75.75 0 0 0-.07 1.058l.662.756H15.75a.75.75 0 0 0 0 1.5h2.862l-.662.756a.75.75 0 0 0 1.13.988l1.75-2a.75.75 0 0 0 0-.988l-1.75-2a.75.75 0 0 0-1.06-.07z" fill="currentColor"/><path d="M5.75 11A2.75 2.75 0 0 0 3 13.75v4a2.75 2.75 0 0 0 2.75 2.75h12.5A2.75 2.75 0 0 0 21 17.75v-4A2.75 2.75 0 0 0 18.25 11H5.75zM6 15.75a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H6.75a.75.75 0 0 1-.75-.75z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}