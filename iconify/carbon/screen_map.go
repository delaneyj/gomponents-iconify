package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 25h8v2h-8zm0-4h8v2h-8zm-3.586-10H23V9h-8v8h2v-4.586L23.586 19L25 17.586L18.414 11z"/><path fill="currentColor" d="M28 3H4c-1.103 0-2 .897-2 2v16c0 1.103.897 2 2 2h8v4H8v2h12v-8H4V5h24v14h2V5c0-1.103-.897-2-2-2ZM18 27h-4v-4h4v4Z"/>`),
		g.Group(children),
	)
}