package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDriveAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 14.44a.62.62 0 0 0 0-.13a.61.61 0 0 1 0-.12l-.05-.12l-6-10.29a1 1 0 0 0-.95-.49H9a1 1 0 0 0-.5.13l-.11.08a.73.73 0 0 0-.09.08a.58.58 0 0 0-.1.12s0 0-.06.08l-6 10.33a1 1 0 0 0 0 1l3 5.08a.83.83 0 0 0 .11.15v.06a1.1 1.1 0 0 0 .44.26a.83.83 0 0 0 .22 0H18a1 1 0 0 0 .86-.49l3-5.14l.05-.12a.61.61 0 0 1 0-.12a.53.53 0 0 0 0-.13a.51.51 0 0 0 0-.13a.59.59 0 0 0 .09-.09ZM6 17.73l-1.79-3.1L9 6.27l.87 1.5l1 1.66L7 15.91Zm6-6.32l1.26 2.16h-2.54Zm5.43 7.3H7.7l1.84-3.14h9.72Zm-1.86-5.14l-4.83-8.28h3.69l4.83 8.28Z"/>`),
		g.Group(children),
	)
}