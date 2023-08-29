package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocameraBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11.5v1c0 3.287 0 4.931.908 6.038a4 4 0 0 0 .554.554C4.57 20 6.212 20 9.5 20c3.287 0 4.931 0 6.038-.908c.202-.166.388-.352.554-.554c.702-.855.861-2.031.897-4.038l.67.33c1.945.972 2.918 1.459 3.63 1.019c.711-.44.711-1.527.711-3.703v-.292c0-2.175 0-3.263-.711-3.703c-.712-.44-1.685.047-3.63 1.02l-.67.329c-.036-2.007-.195-3.183-.897-4.038a4 4 0 0 0-.554-.554C14.43 4 12.788 4 9.5 4c-3.287 0-4.931 0-6.038.908a4 4 0 0 0-.554.554C2 6.57 2 8.212 2 11.5Z"/>`),
		g.Group(children),
	)
}