package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tildes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.056 21.722H3.5v4.556h4.556v-4.556Zm4.555-4.555H8.055v4.555h4.556v-4.555Zm4.556-4.556H12.61v4.556h4.556V12.61Zm4.555 4.556h-4.555v4.555h4.555v-4.555Zm4.556 4.555h-4.556v4.556h4.556v-4.556Zm4.555 4.556h-4.555v4.555h4.555v-4.555Zm4.556 4.555h-4.556v4.556h4.556v-4.556Zm4.555-4.555H35.39v4.555h4.555v-4.555Zm4.556-4.556h-4.556v4.556H44.5v-4.556Z"/>`),
		g.Group(children),
	)
}