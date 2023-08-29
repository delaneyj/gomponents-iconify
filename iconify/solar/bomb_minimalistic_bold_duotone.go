package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BombMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17 14.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Z" opacity=".5"/><path d="M16.671 3.165a.75.75 0 0 0-1.342.67l.5 1a.75.75 0 0 0 1.342-.67l-.5-1Zm3.164 3.665a.75.75 0 0 0-.67 1.341l1 .5a.75.75 0 0 0 .67-1.342l-1-.5Zm.695-2.3a.75.75 0 1 0-1.06-1.06l-1 1a.75.75 0 1 0 1.06 1.06l1-1Zm-5.223 5.224L17.53 7.53a.75.75 0 1 0-1.06-1.06l-2.223 2.223a7.55 7.55 0 0 1 1.06 1.06Z"/></g>`),
		g.Group(children),
	)
}