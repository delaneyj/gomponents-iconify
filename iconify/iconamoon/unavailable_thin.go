package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnavailableThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.5A8.5 8.5 0 0 1 3.5 12h-1a9.5 9.5 0 0 0 9.5 9.5v-1Zm0-17a8.5 8.5 0 0 1 8.5 8.5h1A9.5 9.5 0 0 0 12 2.5v1ZM3.5 12c0-2.347.95-4.472 2.49-6.01l-.708-.708A9.472 9.472 0 0 0 2.5 12h1Zm2.49-6.01A8.472 8.472 0 0 1 12 3.5v-1a9.472 9.472 0 0 0-6.718 2.782l.708.708Zm-.708 0L18.01 18.718l.707-.708L5.99 5.282l-.708.708ZM20.5 12c0 2.347-.95 4.472-2.49 6.01l.707.707A9.472 9.472 0 0 0 21.5 12h-1Zm-2.49 6.01A8.472 8.472 0 0 1 12 20.5v1a9.472 9.472 0 0 0 6.718-2.782l-.708-.708Z"/>`),
		g.Group(children),
	)
}