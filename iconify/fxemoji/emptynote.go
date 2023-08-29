package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emptynote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FFD469" d="M450.812 462.658H74.759a8.802 8.802 0 0 1-8.802-8.802V77.802A8.802 8.802 0 0 1 74.759 69H290.76l168.854 168.854v216.001a8.802 8.802 0 0 1-8.802 8.803z"/><path fill="#FFB636" d="m290.76 69l168.854 168.854H326.651c-19.822 0-35.891-16.069-35.891-35.891V69z"/>`),
		g.Group(children),
	)
}