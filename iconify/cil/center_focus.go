package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 496h120v-32H64a16.019 16.019 0 0 1-16-16V328H16v120a48.054 48.054 0 0 0 48 48ZM48 64a16.019 16.019 0 0 1 16-16h120V16H64a48.054 48.054 0 0 0-48 48v120h32Zm400-48H328v32h120a16.019 16.019 0 0 1 16 16v120h32V64a48.054 48.054 0 0 0-48-48Zm16 432a16.019 16.019 0 0 1-16 16H328v32h120a48.054 48.054 0 0 0 48-48V328h-32Zm-64-192c0-79.4-64.6-144-144-144s-144 64.6-144 144s64.6 144 144 144s144-64.6 144-144ZM256 368a112 112 0 1 1 112-112a112.127 112.127 0 0 1-112 112Z"/>`),
		g.Group(children),
	)
}