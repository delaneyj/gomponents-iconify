package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rabbit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><ellipse cx="24" cy="32" fill="#2F88FF" stroke="#000" stroke-width="4" rx="17" ry="12"/><circle cx="18" cy="29.412" r="2" fill="#fff"/><circle cx="24" cy="35.412" r="2" fill="#fff"/><circle cx="30" cy="29.412" r="2" fill="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12.6672 22C11.3521 18.8333 9.06034 11.1127 10.413 5.91273C10.7887 4.91273 12.2164 3.21273 14.9217 4.41273C15.2974 4.57935 16.1616 5.2126 16.6125 6.4126C17.7397 8.4126 16.0489 21 16.0489 21"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35.3871 22C36.7025 18.8333 38.9324 11.1127 37.5793 5.91273C37.2035 4.91273 35.7754 3.21273 33.0693 4.41273C32.6935 4.57935 31.8291 5.2126 31.378 6.4126C30.2505 8.4126 32.0044 20 32.0044 20"/></g>`),
		g.Group(children),
	)
}