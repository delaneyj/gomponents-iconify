package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jsho(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 23.09v10.82c0 2.719 5.271 5.66 10.154 0m.388-17.846v10.82c0 2.719 5.271 5.66 10.154 0M16.284 22.73a9.25 9.25 0 0 1 2.802 1.914m-4.056-.268a9.25 9.25 0 0 1 2.803 1.914m14.403-15.433v8.822c0 4.182-4.744 4.14-4.744 1.373c0-3.412 5.763-1.706 10.008.374"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.236 14.726a15.187 15.187 0 0 0 4.702-.395"/>`),
		g.Group(children),
	)
}