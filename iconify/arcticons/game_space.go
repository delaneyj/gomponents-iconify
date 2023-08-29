package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameSpace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m23.837 14.422l12.587 7.27l-12.587 7.27l-12.587-7.27l12.587-7.27Z"/><path d="M11.25 21.692V36.23l12.587 7.27l12.587-7.27V21.693M23.837 43.5V28.961"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.002 17.707V8.291m23.399 9.922V4.5m10.597 18.71v-5.503m-24.58-7.215v8.073"/>`),
		g.Group(children),
	)
}