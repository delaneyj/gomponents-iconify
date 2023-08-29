package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarPower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M219.615 401.84h16v88.06h-16v-88.06zm219.94-271.6l21.83-13.18l-21.82-13.16l16-19.8l-25-4.88l8.25-24.12l-25.19 4l-.46-25.49l-22.28 12.29l-9.18-23.8l-16.7 19.24l-16.78-19.24l-9.22 23.8l-22.3-12.31l-.46 25.49l-25.19-4l8.25 24.12l-25 4.88l16 19.8l-21.83 13.18l21.83 13.14h33.41l29.29 76.34l12-13.76l16.74 19.24l9.17-23.76l22.3 12.31l.46-25.49l25.19 4l-8.28-24.18l25-4.88zm-254.55 46.31h-91l31 80.85h91zm108.25 0h-91l30.94 80.85h91zm-70.81 97.42h-91l30.94 80.85h91zm108.25 0h-91l30.94 80.85h91zm73.89 111.87h-262.22L50.615 146.2h262.1zm-19.69-15l-43.31-112.87l-37.28-97.42H70.785l80.47 210.27h233.68z"/>`),
		g.Group(children),
	)
}