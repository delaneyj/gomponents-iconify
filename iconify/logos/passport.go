package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#D6FF00" d="M128 0C57.28 0 0 57.28 0 128h64c0-35.328 28.672-64 64-64V0Z"/><path fill="#34E27A" d="M256 128C256 57.28 198.72 0 128 0v64c35.328 0 64 28.672 64 64h64Z"/><path fill="#00B9F1" d="M128 256c70.72 0 128-57.28 128-128h-64c0 35.328-28.672 64-64 64v64Z"/><path fill="#FFF" d="M64 256V128H0v192h128v-64H64Z"/>`),
		g.Group(children),
	)
}