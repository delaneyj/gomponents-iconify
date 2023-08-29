package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiRouterRoundOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.109 5.684a4.752 4.752 0 0 1 8.781 0a.75.75 0 1 0 1.386-.574a6.252 6.252 0 0 0-11.553 0a.75.75 0 0 0 1.386.574Z"/><path fill-rule="evenodd" d="M7 12.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5ZM6.25 15a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0Z" clip-rule="evenodd"/><path d="M11.25 15a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5H12a.75.75 0 0 1-.75-.75Z"/><path fill-rule="evenodd" d="M7 9.25a5.75 5.75 0 0 0-2.03 11.132l-.64 1.283a.75.75 0 1 0 1.34.67l.805-1.609c.173.016.348.024.525.024h10c.177 0 .352-.008.525-.024l.804 1.61a.75.75 0 1 0 1.342-.671l-.642-1.283A5.752 5.752 0 0 0 17 9.25H7ZM2.75 15A4.25 4.25 0 0 1 7 10.75h10a4.25 4.25 0 0 1 0 8.5H7A4.25 4.25 0 0 1 2.75 15Z" clip-rule="evenodd"/><path d="M16.5 5.25a2.251 2.251 0 0 0-2.16 1.618a.75.75 0 1 1-1.44-.42a3.751 3.751 0 0 1 7.2 0a.75.75 0 1 1-1.44.42A2.251 2.251 0 0 0 16.5 5.25Z"/><path d="M16.5 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`),
		g.Group(children),
	)
}