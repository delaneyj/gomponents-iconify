package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwimmingRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" d="M28 33.1679C30.3069 32.16 32.16 30.3069 33.1679 28"/><path stroke-linecap="round" d="M14.832 28C15.84 30.3069 17.6931 32.16 19.9999 33.1679"/><path stroke-linecap="round" d="M19.9999 14.832C17.6931 15.84 15.84 17.6931 14.832 19.9999"/><path stroke-linecap="round" d="M28 14.832C30.3069 15.84 32.16 17.6931 33.1679 19.9999"/><path stroke-linecap="round" d="M30 40.9758C35.1145 39.1681 39.1681 35.1145 40.9758 30"/><path stroke-linecap="round" d="M7.02441 30C8.83212 35.1145 12.8857 39.1681 18.0002 40.9758"/><path stroke-linecap="round" d="M18.0002 7.02441C12.8857 8.83212 8.83212 12.8857 7.02441 18.0002"/><path stroke-linecap="round" d="M30 7.02441C35.1145 8.83212 39.1681 12.8857 40.9758 18.0002"/><path fill="#2F88FF" stroke-linejoin="round" stroke-miterlimit="2" d="M27 17H21L18 7L20 4H28L30 7L27 17Z"/><path fill="#2F88FF" stroke-linejoin="round" stroke-miterlimit="2" d="M17 21V27L7 30L4 28L4 20L7 18L17 21Z"/><path fill="#2F88FF" stroke-linejoin="round" stroke-miterlimit="2" d="M21 31H27L30 41L28 44H20L18 41L21 31Z"/><path fill="#2F88FF" stroke-linejoin="round" stroke-miterlimit="2" d="M31 27V21L41 18C42.08 18.8 42.92 19.2 44 20V28C42.92 28.8 42.08 29.2 41 30L31 27Z"/></g>`),
		g.Group(children),
	)
}