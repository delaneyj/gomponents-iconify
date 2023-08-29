package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPickerColorColorsDesignDropperEyeEyedropEyedropperPaintingPicker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.46 1.54a3.52 3.52 0 0 0-5 0L2.87 6.13a3 3 0 0 0-.53 3.53L.79 11.21a1 1 0 0 0 0 1.41l.59.59a1 1 0 0 0 1.41 0l1.55-1.55a3 3 0 0 0 3.53-.53l4.59-4.59a3.52 3.52 0 0 0 0-5ZM4.5 1.5l8 8"/>`),
		g.Group(children),
	)
}