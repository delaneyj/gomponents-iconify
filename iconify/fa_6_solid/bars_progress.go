package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarsProgress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 160H320v-32h128v32zM48 64C21.5 64 0 85.5 0 112v64c0 26.5 21.5 48 48 48h416c26.5 0 48-21.5 48-48v-64c0-26.5-21.5-48-48-48H48zm400 288v32H192v-32h256zM48 288c-26.5 0-48 21.5-48 48v64c0 26.5 21.5 48 48 48h416c26.5 0 48-21.5 48-48v-64c0-26.5-21.5-48-48-48H48z"/>`),
		g.Group(children),
	)
}