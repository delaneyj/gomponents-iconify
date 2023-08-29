package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerTowerBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M84 76a12 12 0 0 1 12-12h64a12 12 0 0 1 0 24H96a12 12 0 0 1-12-12Zm12 52h64a12 12 0 0 0 0-24H96a12 12 0 0 0 0 24Zm116-88v176a20 20 0 0 1-20 20H64a20 20 0 0 1-20-20V40a20 20 0 0 1 20-20h128a20 20 0 0 1 20 20Zm-24 4H68v168h120Zm-60 124a16 16 0 1 0 16 16a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}