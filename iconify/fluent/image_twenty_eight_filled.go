package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.75A3.75 3.75 0 0 1 6.75 3h14.5A3.75 3.75 0 0 1 25 6.75v14.5c0 .771-.233 1.488-.632 2.084l-8.97-8.767a2 2 0 0 0-2.796 0l-8.97 8.767A3.733 3.733 0 0 1 3 21.25V6.75Zm1.698 17.64c.59.385 1.295.61 2.052.61h14.5c.757 0 1.462-.225 2.052-.61l-8.952-8.75a.5.5 0 0 0-.7 0l-8.952 8.75ZM20.75 10a2.25 2.25 0 1 0-4.5 0a2.25 2.25 0 0 0 4.5 0Z"/>`),
		g.Group(children),
	)
}