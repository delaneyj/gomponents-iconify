package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 382.1v-88.8c0-14.9 5.9-29.1 16.4-39.6l27.3-27.3l57 57c6.2 6.2 16.4 6.2 22.6 0s6.2-16.4 0-22.6l-57-57l41.4-41.4l57 57c6.2 6.2 16.4 6.2 22.6 0s6.2-16.4 0-22.6l-57-57l41.4-41.4l57 57c6.2 6.2 16.4 6.2 22.6 0s6.2-16.4 0-22.6l-57-57l45.5-45.5C355.2 10.9 381.4 0 408.8 0C465.8 0 512 46.2 512 103.2c0 27.4-10.9 53.6-30.2 73L258.3 399.6a55.924 55.924 0 0 1-39.6 16.4h-88.8L41 505c-9.4 9.4-24.6 9.4-33.9 0s-9.4-24.6 0-33.9l89-89z"/>`),
		g.Group(children),
	)
}