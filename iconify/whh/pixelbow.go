package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixelbow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 960v64h-64v-64h-64v-64h64V608h-64V448h-64v-64h64v64h64v-64h64v192h64v384h-64zM704 320h64v64h-64v-64zm-64 0v-64h64v64h-64zm-64-128H384v-64H128v64H64v-64H0V64h64V0h384v64h192v64h-64v64h64v64h-64v-64zm-384 64h-64v-64h64v64zm64 64h-64v-64h64v64zm64 64h-64v-64h64v64zm64 64h-64v-64h64v64zm64 64h-64v-64h64v64zm64 64h-64v-64h64v64zm64 64h-64v-64h64v64zm64 64h-64v-64h64v64zm64 64h-64v-64h64v64zm64 64h-64v-64h64v64zm64 64h-64v-64h64v64zM704 128v64h-64v-64h64zm0 64h64v64h-64v-64zm128 64v64h-64v-64h64zm64 64v64h-64v-64h64z"/>`),
		g.Group(children),
	)
}