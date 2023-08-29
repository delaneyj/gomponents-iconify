package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommuterBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31.9923 22C31.9974 21.6925 32 21.3833 32 21.0725C32 9.51451 28.4183 4 24 4C19.5817 4 16 9.51451 16 21.0725C16 21.3833 16.0026 21.6925 16.0077 22"/><rect width="22" height="22" x="13" y="22" fill="#2F88FF" rx="2"/></g>`),
		g.Group(children),
	)
}