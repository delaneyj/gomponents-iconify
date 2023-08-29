package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenMapSet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 26h7v2h-7zm0-4h7v2h-7zm-3.586-10H25v-2h-8v8h2v-4.586L25.586 20L27 18.586L20.414 12z"/><path fill="currentColor" d="M7 7h22v12h2V7c0-1.103-.897-2-2-2H7c-1.103 0-2 .897-2 2v15c0 1.103.897 2 2 2h7v4h-4v2h12v-8H7V7Zm13 21h-4v-4h4v4Z"/><path fill="currentColor" d="M26 3V1H3c-1.103 0-2 .897-2 2v15h2V3h23Z"/>`),
		g.Group(children),
	)
}