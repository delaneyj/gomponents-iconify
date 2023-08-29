package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactoryFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 208h-16v-72.98a.76.76 0 0 0 0-.15L209 29.74A16.08 16.08 0 0 0 193.06 16h-18.12a16.08 16.08 0 0 0-15.84 13.74l-11.56 80.91L108.8 81.6A8 8 0 0 0 96 88v32L44.8 81.6A8 8 0 0 0 32 88v120H16a8 8 0 0 0 0 16h224a8 8 0 0 0 0-16Zm-132-24H80a8 8 0 0 1 0-16h28a8 8 0 0 1 0 16Zm68 0h-28a8 8 0 0 1 0-16h28a8 8 0 0 1 0 16Zm-5.33-56l-8.53-6.4l12.8-89.6h18.12l13.72 96Z"/>`),
		g.Group(children),
	)
}