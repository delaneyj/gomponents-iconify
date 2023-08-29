package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraswitchOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17V7h3l1-1h4l1 1h3v10H6Zm2-2h8V9H8v6Zm4-1q.825 0 1.413-.588T14 12q0-.825-.588-1.413T12 10q-.825 0-1.413.588T10 12q0 .825.588 1.413T12 14ZM8.55.5Q9.4.225 10.263.112T12 0q2.35 0 4.438.838t3.7 2.325q1.612 1.487 2.637 3.5T24 11h-2q-.175-1.8-.95-3.363t-1.988-2.75Q17.85 3.7 16.276 2.95t-3.375-.9l1.55 1.55l-1.4 1.4L8.55.5Zm6.9 23q-.85.275-1.712.388T12 24q-2.35 0-4.438-.838t-3.7-2.324q-1.612-1.488-2.637-3.5T0 13h2q.2 1.8.963 3.363t1.974 2.75Q6.15 20.3 7.726 21.05t3.375.9L9.55 20.4l1.4-1.4l4.5 4.5ZM12 12Z"/>`),
		g.Group(children),
	)
}