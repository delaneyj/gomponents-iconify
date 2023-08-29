package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 24L4 20"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 22V18"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 22L40 18"/><circle cx="24" cy="10" r="6" fill="#2F88FF" stroke="#000" stroke-width="4"/><ellipse cx="24" cy="30" fill="#2F88FF" stroke="#000" stroke-width="4" rx="13" ry="14"/><circle cx="24" cy="26" r="2" fill="#fff"/><circle cx="24" cy="31" r="2" fill="#fff"/><circle cx="24" cy="36" r="2" fill="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 20L36 24"/></g>`),
		g.Group(children),
	)
}