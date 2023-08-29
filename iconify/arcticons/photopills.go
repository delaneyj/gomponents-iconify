package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photopills(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h18.555v18.555h0H5.5h0V5.5h0zm18.555 0v18.555h9.83a8.597 8.597 0 0 0 8.615-8.616v-1.322A8.597 8.597 0 0 0 33.885 5.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24.055v9.83a8.597 8.597 0 0 0 8.616 8.615h1.322a8.597 8.597 0 0 0 8.616-8.615v-9.83Z"/>`),
		g.Group(children),
	)
}