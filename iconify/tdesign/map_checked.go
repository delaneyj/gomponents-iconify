package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V12h-2V5.5l-4 1.667V12h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926l7-4.084ZM8 16.833V4.741L4 7.074V18.5l4-1.667Zm15.657-2.247l-7.07 7.071l-4.243-4.243L13.758 16l2.828 2.829l5.657-5.657l1.414 1.414Z"/>`),
		g.Group(children),
	)
}