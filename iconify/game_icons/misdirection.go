package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Misdirection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M273.5 22.48L166 86.15l45.2 16.05l-42.8 120.6l60.2 21.4l42.9-120.6l45.2 16.1zM62.4 35.28l-4.77 59.96l-22.59-1.91l34 48.07L110 99.44l-22.46-1.8l4.83-60.11zm329 58.44l-30.1 27.98l55.8 60.2l-22.7 20.9l79.6 10.3l-4.2-80l-22.6 20.9zM45.51 202.5L32.3 237l68.6 26.7l-10.22 25.8l68.42-22.7l-34.6-63.2l-10.2 25.9zm377.79 51.8l-94.9 7.6l24.7 91.6l21-29.8l79.4 56.3l28-39.8l-79.4-56.1zm-180.7 14.5L106.8 370.7l-38.27-51.3l-38.05 161.7l165.32 8.4l-38.1-51l136.1-101.8zm92.8 115.6l-22.3 15.8l31.3 44.9l-16.9 11.7l51.6 14.3l4.8-53.3l-17 11.7z"/>`),
		g.Group(children),
	)
}