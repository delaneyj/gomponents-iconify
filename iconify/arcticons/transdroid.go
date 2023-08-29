package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43 25.43l3.13-1.81L43 21.8l2.71-2.37l-3.41-1.18l2.19-2.86l-3.59-.49l1.58-3.23l-3.61.21l.92-3.47l-3.5.91l.22-3.59l-3.25 1.57l-.49-3.56l-2.88 2.16l-1.18-3.39l-2.4 2.68l-1.82-3.1l-1.83 3.1l-2.4-2.69l-1.18 3.39l-2.88-2.17l-.5 3.55l-3.25-1.57l.21 3.58l-3.5-.91l.92 3.47l-3.61-.22L8 14.85l-3.58.49l2.22 2.86l-3.42 1.17l2.7 2.38l-3.12 1.81l3.12 1.82l-2.71 2.38l3.42 1.17l-2.19 2.86l3.58.5l-1.59 3.22L10 35.3l-.92 3.47l3.5-.9l-.22 3.58l3.25-1.56l.49 3.55l2.9-2.16l1.17 3.39L22.61 42l1.82 3.1l1.83-3.1l2.39 2.69l1.19-3.39l2.88 2.18l.5-3.56l3.24 1.57l-.21-3.58l3.5.91l-.91-3.47l3.61.22l-1.58-3.23l3.58-.48L42.27 29l3.42-1.17Zm-8.73-1.89a10 10 0 0 1-10 10h0a10 10 0 0 1-3.88-19.25v8.56H18l6.34 6.68l6.33-6.68h-2.18v-8.4a10 10 0 0 1 5.77 9.07Z"/>`),
		g.Group(children),
	)
}