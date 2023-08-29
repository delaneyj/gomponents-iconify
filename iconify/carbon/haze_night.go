package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HazeNight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 22H13a1 1 0 0 0 0 2h9.677a11.1 11.1 0 0 1-7.556 2.998c-.138 0-.278.004-.418 0a11.095 11.095 0 0 1-3.2-21.584a15.153 15.153 0 0 0-.05 6.586H13.5a13.15 13.15 0 0 1 .39-7.62A1.015 1.015 0 0 0 12.98 3a.983.983 0 0 0-.175.016a13.096 13.096 0 0 0 1.825 25.981c.164.006.328 0 .49 0A13.04 13.04 0 0 0 25.379 24H29a1 1 0 0 0 0-2Z"/><path fill="currentColor" d="M27 20H17a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm-2-4H13a1 1 0 0 1 0-2h12a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}