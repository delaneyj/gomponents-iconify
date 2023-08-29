package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gmail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.585 768h-896q-26 0-45-19t-19-45V64q0-27 18.5-45.5T64.585 0h896q27 0 45.5 18.5t18.5 45.5v640q0 26-18.5 45t-45.5 19zm0-672q0-13-9.5-22.5t-22.5-9.5h-64q-20 0-31.5 1.5t-30 11t-37.5 28.5l-253 253l-253-253q-19-19-37.5-28.5t-29.5-11t-32-1.5h-64q-13 0-22.5 9.5t-9.5 22.5v576q0 13 9.5 22.5t22.5 9.5h64q13 0 22.5-9.5t9.5-22.5v-33l211-210l74 74q15 15 35 7q20 8 35-7l74-74l211 210v33q0 13 9.5 22.5t22.5 9.5h64q13 0 22.5-9.5t9.5-22.5V96zm-128 453l-166-166l166-165v331zm-640-331l166 165l-166 166V218z"/>`),
		g.Group(children),
	)
}