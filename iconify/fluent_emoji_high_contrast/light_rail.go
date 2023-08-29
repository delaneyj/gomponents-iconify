package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightRail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M11 20a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2Z"/><path d="M13.706 11.078A8 8 0 0 1 20.68 7H30v23H2v-3h6.78a1 1 0 0 1-.97-.758l-.327-1.312a4.999 4.999 0 0 1 .492-3.664l5.731-10.188ZM10.96 18l-2.113 3.756A4 4 0 0 0 8.398 23H29v-1h-7a1 1 0 1 1 0-2h7v-2h-5.554c-1.535 0-2.498-1.659-1.736-2.992l1.138-1.993A4 4 0 0 1 26.321 11H29V9h-3.324a5 5 0 0 0-4.16 2.226l-3.032 4.547A5 5 0 0 1 14.324 18H10.96Zm6.113-9a6.996 6.996 0 0 0-2.495 2.568L12.085 16h.342a5 5 0 0 0 4.069-2.094l2.375-3.325A1 1 0 0 0 18.056 9h-.984Z"/></g>`),
		g.Group(children),
	)
}