package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V11h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926l7-4.084ZM8 16.833V4.741L4 7.074V18.5l4-1.667ZM18 14c-.69 0-1.25.56-1.25 1.25V16h5.749v6.5h-9V16h1.251v-.75a3.25 3.25 0 0 1 5.541-2.305l.71.705l-1.41 1.418l-.71-.705A1.243 1.243 0 0 0 18 14Zm-2.501 4v2.5h5V18h-5Z"/>`),
		g.Group(children),
	)
}