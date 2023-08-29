package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnifeFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 4V44"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 5V15C8 20 14 20 14 20C14 20 20 20 20 15V5"/><path fill="#2F88FF" d="M30 12C30 4 38 4 38 4V21H30V12Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 21H30V12C30 4 38 4 38 4V21ZM38 21V44"/></g>`),
		g.Group(children),
	)
}