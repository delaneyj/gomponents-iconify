package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseScrollWheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 5l.53-.53a.75.75 0 0 0-1.06 0L12 5Zm0 8l-.53.53a.75.75 0 0 0 1.06 0L12 13ZM9.47 6.47a.75.75 0 0 0 1.06 1.06L9.47 6.47Zm4 1.06a.75.75 0 1 0 1.06-1.06l-1.06 1.06Zm-2.94 2.94a.75.75 0 1 0-1.06 1.06l1.06-1.06Zm4 1.06a.75.75 0 1 0-1.06-1.06l1.06 1.06ZM3.25 10v4h1.5v-4h-1.5Zm17.5 4v-4h-1.5v4h1.5Zm-9.5-9v8h1.5V5h-1.5Zm.22-.53l-2 2l1.06 1.06l2-2l-1.06-1.06Zm0 1.06l2 2l1.06-1.06l-2-2l-1.06 1.06Zm1.06 6.94l-2-2l-1.06 1.06l2 2l1.06-1.06Zm0 1.06l2-2l-1.06-1.06l-2 2l1.06 1.06ZM20.75 10A8.75 8.75 0 0 0 12 1.25v1.5A7.25 7.25 0 0 1 19.25 10h1.5ZM12 22.75A8.75 8.75 0 0 0 20.75 14h-1.5A7.25 7.25 0 0 1 12 21.25v1.5ZM3.25 14A8.75 8.75 0 0 0 12 22.75v-1.5A7.25 7.25 0 0 1 4.75 14h-1.5Zm1.5-4A7.25 7.25 0 0 1 12 2.75v-1.5A8.75 8.75 0 0 0 3.25 10h1.5Z"/>`),
		g.Group(children),
	)
}