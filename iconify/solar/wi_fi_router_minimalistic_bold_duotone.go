package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiRouterMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M16.5 4.5a4.752 4.752 0 0 0-4.39 2.934a.75.75 0 1 1-1.387-.574a6.252 6.252 0 0 1 11.553 0a.75.75 0 0 1-1.386.574A4.752 4.752 0 0 0 16.5 4.5Z" clip-rule="evenodd" opacity=".4"/><path fill-rule="evenodd" d="M16.5 7a2.251 2.251 0 0 0-2.16 1.618a.75.75 0 0 1-1.44-.42a3.751 3.751 0 0 1 7.2 0a.75.75 0 1 1-1.44.42A2.251 2.251 0 0 0 16.5 7Z" clip-rule="evenodd" opacity=".7"/><path fill-rule="evenodd" d="M2 15.75c0-1.886 0-2.828.586-3.414c.586-.586 1.528-.586 3.414-.586h12c1.886 0 2.828 0 3.414.586c.586.586.586 1.528.586 3.414c0 1.886 0 2.828-.586 3.414c-.586.586-1.528.586-3.414.586H6c-1.886 0-2.828 0-3.414-.586C2 18.578 2 17.636 2 15.75Zm4 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/><path d="M17.25 8.75a.75.75 0 0 0-1.5 0v3h1.5v-3Z"/></g>`),
		g.Group(children),
	)
}