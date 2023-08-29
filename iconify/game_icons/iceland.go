package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iceland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m152 238.5l19.7-60l19.3 21.6l7.4-72.4l50.8 51.4l-18.6-43.5l65.5-11.9l22.7 23.8c21.4-17.4 32.1-31.6 40.1-59.39l39.5-6.82l4 30.61l91.1 84.8c-5.3 49.1-9.1 98.8-58.3 133.4L311.5 390L282 430.7c-59.5-1.5-97.8-25.2-128.9-56.6H71.68l-7.91-26l52.53-8.4l-30.48-67.9H19.11l3.38-19.8l100.61-3.9l-26.44-18.9c-.93-4.3 22.84-18.7 21.54-22.2c-11-28.2-49.42-13.5-99.66-5.3c6.51-34.3 32.09-81 65.01-113.06c30.65 15.26 47.05 36.86 68.95 55.96c-9.4 36-5 57.8-.5 93.9z"/>`),
		g.Group(children),
	)
}