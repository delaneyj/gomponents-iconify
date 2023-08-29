package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 960v-64H192v64H0V0h192v64h640V0h192v960H832zM128 64H64v64h64V64zm0 128H64v64h64v-64zm0 128H64v64h64v-64zm0 128H64v64h64v-64zm0 128H64v64h64v-64zm0 128H64v64h64v-64zm0 128H64v64h64v-64zm704-704H192v704h640V128zm128-64h-64v64h64V64zm0 128h-64v64h64v-64zm0 128h-64v64h64v-64zm0 128h-64v64h64v-64zm0 128h-64v64h64v-64zm0 128h-64v64h64v-64zm0 128h-64v64h64v-64z"/>`),
		g.Group(children),
	)
}