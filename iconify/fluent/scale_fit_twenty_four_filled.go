package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleFitTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.25 4A2.25 2.25 0 0 0 2 6.25v11.5A2.25 2.25 0 0 0 4.25 20h15.5A2.25 2.25 0 0 0 22 17.75V6.25A2.25 2.25 0 0 0 19.75 4H4.25Zm3.744 5.436a.75.75 0 0 1 .07 1.058l-.661.756h2.862a.75.75 0 0 1 0 1.5H7.403l.661.756a.75.75 0 0 1-1.128.988l-1.75-2a.75.75 0 0 1 0-.988l1.75-2a.75.75 0 0 1 1.058-.07Zm7.956 1.058a.75.75 0 0 1 1.13-.988l1.75 2a.75.75 0 0 1 0 .988l-1.75 2a.75.75 0 0 1-1.13-.988l.662-.756H13.75a.75.75 0 0 1 0-1.5h2.862l-.662-.756Z"/>`),
		g.Group(children),
	)
}