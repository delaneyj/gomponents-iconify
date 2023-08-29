package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buienalarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.824 32.597a3.11 3.11 0 1 0-2.816-4.412a5.089 5.089 0 0 0-9.736-1.665a4.824 4.824 0 1 0-9.398-2.164a4.147 4.147 0 1 0-.654 8.241ZM24.406 41.96c.421-1.2-.026-6.873-.026-6.873s-3.896 4.149-4.317 5.348a2.302 2.302 0 0 0 4.343 1.525Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.868 38.647c.37-1.056-.024-6.05-.024-6.05s-3.43 3.652-3.8 4.708a2.026 2.026 0 0 0 3.824 1.342Zm-14.374 0c.37-1.056-.024-6.05-.024-6.05s-3.43 3.652-3.8 4.708a2.026 2.026 0 0 0 3.824 1.342ZM32.431 4.5L21.355 18.006l17.234 2.839L32.431 4.5zm-1.63 9.891l.75-4.551"/><circle cx="30.422" cy="16.696" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}