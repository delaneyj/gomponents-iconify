package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SouthAfricaFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M18 18v104.324L222.445 256L18 389.676V494h55.553l273.765-179H494V197H347.318L73.553 18zm88.447 0l246.235 161H494V18zm246.235 315L106.447 494H494V333z"/>`),
		g.Group(children),
	)
}