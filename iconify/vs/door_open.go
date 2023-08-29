package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M400 200H0v1600h400v200l800-200V200L400 0v200zM100 300h300v1400H100V300zm479 779q-24 0-43.5-25T516 995t19.5-59.5T579 910t43 25t19 60q0 34-19 59t-43 25z"/>`),
		g.Group(children),
	)
}