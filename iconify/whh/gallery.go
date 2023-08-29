package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M128 1024V128h896v896H128zm832-832H192v640h768V192zm-576 64q53 0 90.5 37.5T512 384t-37.5 90.5T384 512t-90-37.5t-38-90.5v-1q1-53 38-90t90-37zm47 337l174 175H256v-93l82-82q18-18 46.5-18t46.5 18zm259-128q18-18 46.5-18t46.5 18l113 113v190H695L542 615zM64 896H0V0h896v64H64v832z"/>`),
		g.Group(children),
	)
}