package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiRouterRoundBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16.5 3.5a4.752 4.752 0 0 0-4.39 2.934a.75.75 0 1 1-1.387-.574a6.252 6.252 0 0 1 11.553 0a.75.75 0 0 1-1.385.574A4.752 4.752 0 0 0 16.5 3.5Z"/><path d="M16.5 6a2.251 2.251 0 0 0-2.16 1.618a.75.75 0 0 1-1.44-.42a3.751 3.751 0 0 1 7.2 0a.75.75 0 1 1-1.44.42A2.251 2.251 0 0 0 16.5 6Z"/><path d="M16.5 8.75a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM7 14a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Z"/><path fill-rule="evenodd" d="M2 14.75a5 5 0 0 1 5-5h10a5 5 0 0 1 1.531 9.761l1.112 1.853a.75.75 0 0 1-1.286.772l-1.432-2.386h-9.85l-1.432 2.386a.75.75 0 0 1-1.286-.772l1.112-1.853A5.002 5.002 0 0 1 2 14.75Zm2.75 0a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0ZM12 14a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5H12Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}