package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V11h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926l7-4.084ZM8 16.833V4.741L4 7.074V18.5l4-1.667Zm7.172-4.076L18 15.586l2.829-2.829l1.414 1.415L19.415 17l2.828 2.828l-1.414 1.415L18 18.414l-2.828 2.829l-1.414-1.415L16.586 17l-2.828-2.828l1.414-1.415Z"/>`),
		g.Group(children),
	)
}