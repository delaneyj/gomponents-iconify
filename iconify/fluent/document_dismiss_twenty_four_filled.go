package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDismissTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2v6a2 2 0 0 0 2 2h6v10a2 2 0 0 1-2 2h-6.81A6.5 6.5 0 0 0 4 11.498V4a2 2 0 0 1 2-2h6Zm1.5.5V8a.5.5 0 0 0 .5.5h5.5l-6-6Zm-7 9.5a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11Zm2.478 3.731l-1.77 1.77l1.77 1.769a.5.5 0 1 1-.707.707l-1.77-1.77l-1.769 1.768a.5.5 0 1 1-.707-.708L5.794 17.5l-1.769-1.77a.5.5 0 1 1 .707-.707l1.77 1.769l1.77-1.769a.5.5 0 0 1 .706.707Z"/>`),
		g.Group(children),
	)
}