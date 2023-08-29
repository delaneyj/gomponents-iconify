package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuddiconsFriends(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.75 5.77C8.75 4.39 7 2 7 2S5.25 4.39 5.25 5.77S5.9 7.5 7 7.5s1.75-.35 1.75-1.73zm6 0C14.75 4.39 13 2 13 2s-1.75 2.39-1.75 3.77S11.9 7.5 13 7.5s1.75-.35 1.75-1.73zM9 17V9c0-.55-.45-1-1-1H6c-.55 0-1 .45-1 1v8c0 .55.45 1 1 1h2c.55 0 1-.45 1-1zm6 0V9c0-.55-.45-1-1-1h-2c-.55 0-1 .45-1 1v8c0 .55.45 1 1 1h2c.55 0 1-.45 1-1zm-9-6l2-1v2l-2 1v-2zm6 0l2-1v2l-2 1v-2zm-6 3l2-1v2l-2 1v-2zm6 0l2-1v2l-2 1v-2z"/>`),
		g.Group(children),
	)
}