package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedroomParentSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17h1.5v-1.5h11V17H19v-6.5h-.75V7H5.75v3.5H5V17Zm7.75-6.5v-2h4v2h-4Zm-5.5 0v-2h4v2h-4ZM6.5 14v-2h11v2h-11ZM2 22V2h20v20H2Z"/>`),
		g.Group(children),
	)
}