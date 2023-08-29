package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sobriety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.22 19.008l-1.272.622l-.645 1.26l-.622-1.271l-1.26-.645l1.272-.623l.645-1.26l.622 1.272Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.175 12.591c-9.205.475-15.744 3.052-15.687 6.117c.112 5.977 10.267 5.77 11.024 9.33c.773 3.633-4.655 7.238-10.121 7.355c-6.056.13-9.566-.598-9.566-2.28s4.239-4.43 13.313-4.946"/>`),
		g.Group(children),
	)
}