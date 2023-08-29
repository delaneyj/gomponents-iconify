package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M17.873 17.254v476.558h176.373L315.14 355.4l-137.83 63.075l70.667-108.043l-31.43 16.517c13.426-76.567 48.536-132.702 102.05-197.208l-11.85 112.803l71.17-140.2l-6.72 151.587L470.27 77.254c12.35-17.27 20.207-38 24.748-60H17.873zm216.71 55.97L213.02 179.57l73.123-103.343l-21.637 93.414c-35.604 51.076-59.427 102.66-68.56 168.135l-62.436 32.81l63.072-89.355l-117.97 56.065l-37.962-65.992c10.632-76.265 43.808-139.937 99.284-191.56l94.648-6.518z"/>`),
		g.Group(children),
	)
}