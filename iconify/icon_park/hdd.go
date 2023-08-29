package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="30" height="40" x="9" y="4" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><circle cx="32" cy="10" r="2" fill="#fff"/><circle cx="16" cy="10" r="2" fill="#fff"/><circle cx="32" cy="38" r="2" fill="#fff"/><circle cx="16" cy="38" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 30C28.4183 30 32 26.4183 32 22C32 17.5817 28.4183 14 24 14C19.5817 14 16 17.5817 16 22"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 22L18 28"/></g>`),
		g.Group(children),
	)
}