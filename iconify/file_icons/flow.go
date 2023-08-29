package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M342.066 440.933v-29.62l14.732 14.851V283.013h-29.682v-71.166l53.326 53.759V122.455H259.879L138.408 0v52.333H29.391l28.498 28.729H0l141.804 143.582H29.751l112.489 113.4H0l171.714 173.185V143.151h87.258l35.931 36.223l-80.111-.004v143.151h39.196l45.959 46.331h-89.173L210.77 512h201.792z"/>`),
		g.Group(children),
	)
}