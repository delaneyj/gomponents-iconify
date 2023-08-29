package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cubes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m6.5 10.5l5.5 3l5.5-3v-6l-5.5-3l-5.5 3v6Zm0-6l5.5 3l5.5-3m-5.5 3v6v-6Zm-11 12l5.5 3l5.5-3v-6l-5.5-3l-5.5 3v6Zm0-6l5.5 3l5.5-3m-5.5 3v6v-6Zm5.5 3l5.5 3l5.5-3v-6l-5.5-3l-5.5 3v6Zm0-6l5.5 3l5.5-3m-5.5 3v6v-6Z"/>`),
		g.Group(children),
	)
}