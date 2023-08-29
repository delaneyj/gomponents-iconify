package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AakashTutor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.994" cy="8.695" r="3.196" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.963 42.5a3.166 3.166 0 0 0 2.943-2.018c.673-1.598-.084-3.448-1.766-4.12c-7.148-2.86-11.352-8.662-13.118-17.576c-.337-1.682-2.019-2.859-3.7-2.522s-2.86 2.018-2.523 3.7c2.102 11.016 7.82 18.5 16.902 22.284c.42.168.841.252 1.262.252Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.755 19.668c-1.79 8.459-5.983 13.941-12.906 16.693c-1.598.673-2.439 2.523-1.766 4.12a3.166 3.166 0 0 0 2.943 2.02c.42 0 .841-.085 1.178-.253c6.222-2.523 10.763-6.727 13.79-12.698"/>`),
		g.Group(children),
	)
}