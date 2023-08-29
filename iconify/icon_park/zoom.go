package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M20 43C15.2742 41.2327 11.3325 35.1851 9.35971 31.6428C8.50794 30.1134 8.95664 28.2347 10.3236 27.1411C11.8473 25.9222 14.0438 26.0438 15.4236 27.4236L17 29V20.5C17 19.1193 18.1193 18 19.5 18C20.8807 18 22 19.1193 22 20.5V16.5C22 15.1193 23.1193 14 24.5 14C25.8807 14 27 15.1193 27 16.5V24.5C27 23.1193 28.1193 22 29.5 22C30.8807 22 32 23.1193 32 24.5V27.5C32 26.1193 33.1193 25 34.5 25C35.8807 25 37 26.1193 37 27.5V35.368C37 36.4383 36.7354 37.496 36.1185 38.3707C35.0949 39.8219 33.255 42.0336 31 43C27.5 44.5 24.3701 44.6343 20 43Z"/><path d="M13 8L35 8"/><path d="M17.0003 12L13 8L17 4"/><path d="M31 4L35 8L31 12"/></g>`),
		g.Group(children),
	)
}