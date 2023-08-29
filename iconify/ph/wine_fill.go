package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m205.33 95.67l-21.77-73.93a8 8 0 0 0-7.67-5.74H80.11a8 8 0 0 0-7.67 5.74L50.67 95.67a63.46 63.46 0 0 0 17.42 64.67A87.41 87.41 0 0 0 120 183.63V224H88a8 8 0 1 0 0 16h80a8 8 0 1 0 0-16h-32v-40.37a87.39 87.39 0 0 0 51.91-23.29a63.48 63.48 0 0 0 17.42-64.67ZM86.09 32h83.82L190 100.19c.09.3.17.6.25.9c-21.42 7.68-45.54-1.6-58.63-8.23c-25.19-12.75-45.19-14.37-59.94-11.93Z"/>`),
		g.Group(children),
	)
}