package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SofaThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11 18h9a2 2 0 1 0 0-4H4a2 2 0 1 0 0 4h3m-2.5-4l-.075-.299c-1.087-4.347-1.63-6.52-.56-8.023c.068-.095.14-.186.216-.275C5.278 4 7.519 4 12 4c.723 0 1.388 0 2 .006M19.5 14l.075-.299c1.086-4.347 1.63-6.52.559-8.023a4.002 4.002 0 0 0-.215-.275c-.462-.54-1.078-.873-1.919-1.078M20 20v-2M4 20v-2"/>`),
		g.Group(children),
	)
}