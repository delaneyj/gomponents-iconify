package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiRouterRoundBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7 14a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Z"/><path fill-rule="evenodd" d="M2 14.75a5 5 0 0 1 5-5h10a5 5 0 0 1 0 10H7a5 5 0 0 1-5-5Zm2.75 0a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0ZM12 14a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5H12Z" clip-rule="evenodd"/><path d="m5.469 19.511l-1.112 1.853a.75.75 0 0 0 1.286.772l1.432-2.386H7a4.997 4.997 0 0 1-1.531-.239Zm11.456.239l1.432 2.386a.75.75 0 0 0 1.286-.772l-1.112-1.853a4.998 4.998 0 0 1-1.531.24h-.075Z" opacity=".5"/><path fill-rule="evenodd" d="M16.5 3.5a4.752 4.752 0 0 0-4.39 2.934a.75.75 0 1 1-1.387-.574a6.252 6.252 0 0 1 11.553 0a.75.75 0 0 1-1.386.574A4.752 4.752 0 0 0 16.5 3.5Z" clip-rule="evenodd" opacity=".4"/><path fill-rule="evenodd" d="M16.5 6a2.251 2.251 0 0 0-2.16 1.618a.75.75 0 0 1-1.44-.42a3.751 3.751 0 0 1 7.2 0a.75.75 0 1 1-1.44.42A2.251 2.251 0 0 0 16.5 6Z" clip-rule="evenodd" opacity=".7"/><path d="M17.5 7.75a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		g.Group(children),
	)
}