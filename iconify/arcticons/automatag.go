package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Automatag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 6.577h-9.769v25.035"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.73 31.612a6.465 6.465 0 0 1-6.464 6.465h0a6.465 6.465 0 0 1-6.466-6.465h0a6.465 6.465 0 0 1 6.465-6.465h0a6.465 6.465 0 0 1 6.466 6.465Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.001 41.267a9.811 9.811 0 1 1 7.897-7.792m-16.57-8.801L4.5 8.846"/>`),
		g.Group(children),
	)
}