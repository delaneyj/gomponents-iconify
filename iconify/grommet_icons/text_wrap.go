package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M17 10h7h-7ZM1 14h13V2H1v12Zm5-8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m11 0h7h-7Zm0-4h7h-7Zm0 12h7h-7ZM0 18h24H0Zm0 4h24H0Zm14-8v-1l-4-5l-3 3l-1-1l-4 4h12Z"/>`),
		g.Group(children),
	)
}