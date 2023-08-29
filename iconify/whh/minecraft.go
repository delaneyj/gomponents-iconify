package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minecraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024L0 768V256L512 0l512 256v512zm448-736L512 64L64 288v96h64v64h64v-64h64v64h64v64h64v-64l64 32v96h64v-64h64v64h64v-64h64v-96l64-32v64h64v-96l64-32v64h64v-96zm-704 96z"/>`),
		g.Group(children),
	)
}