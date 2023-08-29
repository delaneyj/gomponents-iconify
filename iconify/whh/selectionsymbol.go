package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Selectionsymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M453 709q-14 15-32 25q27 46 27 98q0 80-56 136t-136 56t-136-56t-56-136q0-22 7-48l121 48q0 26 18.5 45t45 19t45.5-18.5t19-45.5q0-28-16.5-45t-52-38t-59.5-45L64 576q-31-31-47.5-63.5T0 448q0-77 58-134q14-14 32-26q-26-44-26-96q0-80 56-136T256 0t136 56t56 136q0 22-7 48l-121-48q0-27-19-45.5T255.5 128t-45 18.5T192 192q0 28 16.5 45t52.5 38.5t59 44.5l128 128q35 35 49.5 61.5T512 576q0 74-59 133zm-89-183L241 403q-20-20-47.5-20t-47 19.5t-19.5 47t20 47.5l123 123q20 20 47.5 20t46.5-19.5t19-47t-19-47.5z"/>`),
		g.Group(children),
	)
}