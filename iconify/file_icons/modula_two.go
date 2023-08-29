package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModulaTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 325H188V0h324v325zM151 0H0v325h151V0zM0 512h151V362H0v150zm188 0h324V362H188v150zm-37-187v37h37v-37h-37z"/>`),
		g.Group(children),
	)
}