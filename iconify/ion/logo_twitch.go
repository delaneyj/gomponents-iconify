package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoTwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m80 32l-32 80v304h96v64h64l64-64h80l112-112V32Zm336 256l-64 64h-96l-64 64v-64h-80V80h304Z"/><path fill="currentColor" d="M320 143h48v129h-48zm-112 0h48v129h-48z"/>`),
		g.Group(children),
	)
}