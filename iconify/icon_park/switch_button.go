package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M36 4H12C7.58172 4 4 7.58172 4 12C4 16.4183 7.58172 20 12 20H36C40.4183 20 44 16.4183 44 12C44 7.58172 40.4183 4 36 4Z"/><path fill="#2F88FF" stroke="#000" d="M36 28H12C7.58172 28 4 31.5817 4 36C4 40.4183 7.58172 44 12 44H36C40.4183 44 44 40.4183 44 36C44 31.5817 40.4183 28 36 28Z"/><path fill="#43CCF8" stroke="#fff" d="M36 14C37.1046 14 38 13.1046 38 12C38 10.8954 37.1046 10 36 10C34.8954 10 34 10.8954 34 12C34 13.1046 34.8954 14 36 14Z"/><path fill="#43CCF8" stroke="#fff" d="M12 38C13.1046 38 14 37.1046 14 36C14 34.8954 13.1046 34 12 34C10.8954 34 10 34.8954 10 36C10 37.1046 10.8954 38 12 38Z"/></g>`),
		g.Group(children),
	)
}