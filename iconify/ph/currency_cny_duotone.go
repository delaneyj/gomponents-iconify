package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyCnyDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 64v56H64V64Z" opacity=".2"/><path d="M56 64a8 8 0 0 1 8-8h128a8 8 0 0 1 0 16H64a8 8 0 0 1-8-8Zm160 104a8 8 0 0 0-8 8v16h-32a16 16 0 0 1-16-16v-48h48a8 8 0 0 0 0-16H48a8 8 0 0 0 0 16h48v8a56.06 56.06 0 0 1-56 56a8 8 0 0 0 0 16a72.08 72.08 0 0 0 72-72v-8h32v48a32 32 0 0 0 32 32h40a8 8 0 0 0 8-8v-24a8 8 0 0 0-8-8Z"/></g>`),
		g.Group(children),
	)
}