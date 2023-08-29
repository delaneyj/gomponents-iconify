package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SofaThreeLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M4 18h16a2 2 0 1 0 0-4H4a2 2 0 1 0 0 4Z"/><path d="m4.5 14l-.075-.299c-1.087-4.347-1.63-6.52-.56-8.023c.068-.095.14-.186.216-.275C5.278 4 7.519 4 12 4s6.722 0 7.92 1.403c.075.089.146.18.214.275c1.07 1.502.527 3.676-.56 8.023L19.5 14" opacity=".5"/><path stroke-linecap="round" d="M20 20v-2M4 20v-2" opacity=".5"/></g>`),
		g.Group(children),
	)
}