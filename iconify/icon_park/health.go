package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Health(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M39 6H9C7.34315 6 6 7.34315 6 9V39C6 40.6569 7.34315 42 9 42H39C40.6569 42 42 40.6569 42 39V9C42 7.34315 40.6569 6 39 6Z"/><path fill="#43CCF8" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M27.3 12C25.4775 12 24 13.4347 24 15.2045C24 18.4091 27.9 21.3223 30 22C32.1 21.3223 36 18.4091 36 15.2045C36 13.4347 34.5225 12 32.7 12C31.5839 12 30.5972 12.538 30 13.3616C29.4028 12.538 28.4161 12 27.3 12Z"/></g>`),
		g.Group(children),
	)
}