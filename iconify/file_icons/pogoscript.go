package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pogoscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M285.14 0v28.945h-95.388V18.617h-25.74v329.656h-42.883V18.617H95.387v10.328H0V0h95.387v7.897h94.365V0h95.388zm-94.864 358.444H94.864v-20.37H5.884v21.442h17.153v20.369h105.06v99.954h-8.307V512h45.56v-32.161h-8.307v-99.954h105.06v-20.37h17.153v-21.44h-88.98v20.369z"/>`),
		g.Group(children),
	)
}