package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushPinSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.93 210.62a8 8 0 1 1-11.84 10.76l-21.84-24a89.75 89.75 0 0 1-11.58 20.27A15.93 15.93 0 0 1 157 224h-1.13a16 16 0 0 1-11.32-4.69L96.29 171l-42.63 42.66a8 8 0 0 1-11.32-11.32L85 159.71l-48.3-48.3A16 16 0 0 1 38 87.63A69.82 69.82 0 0 1 67.37 73.2L42.08 45.38a8 8 0 1 1 11.84-10.76L182.3 175.83Zm21.39-129.25l-60.69-60.68a16 16 0 0 0-22.63 0l-41.23 41.37a8 8 0 0 0-.25 11l68.57 75.43a8 8 0 0 0 5.74 2.62h.17a8 8 0 0 0 5.66-2.35L235.32 104a16 16 0 0 0 0-22.63Z"/>`),
		g.Group(children),
	)
}