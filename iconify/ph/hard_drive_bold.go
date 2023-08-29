package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 60H32a20 20 0 0 0-20 20v96a20 20 0 0 0 20 20h192a20 20 0 0 0 20-20V80a20 20 0 0 0-20-20Zm-4 112H36V84h184Zm-56-44a16 16 0 1 1 16 16a16 16 0 0 1-16-16Z"/>`),
		g.Group(children),
	)
}