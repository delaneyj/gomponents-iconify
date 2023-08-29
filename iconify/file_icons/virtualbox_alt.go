package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualboxAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0v512h511.308V0H0zm471.177 40.13v252.69h-58.513L338.266 69.666l-54.071 147.912l-50.92-147.163l-65.702 239.958l-41.128-131.31H40.131V40.13h431.046zM40.131 471.87V219.193h56.83l73.198 233.7l67.584-246.837l45.415 131.246l53.365-145.981l47.218 141.629h87.436v138.919H40.131z"/>`),
		g.Group(children),
	)
}