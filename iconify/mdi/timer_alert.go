package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.45 5.97c-.45-.51-.9-.97-1.41-1.41L15.62 6c-1.55-1.26-3.5-2-5.62-2a9 9 0 0 0 0 18c5 0 9-4.03 9-9c0-2.12-.74-4.07-1.97-5.61l1.42-1.42M11 14H9V7h2v7m2-11H7V1h6v2m10 4v6h-2V7h2m-2 8h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}