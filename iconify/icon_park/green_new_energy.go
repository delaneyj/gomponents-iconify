package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenNewEnergy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42.3357 16C39.2495 8.93638 32.2012 4 23.9999 4C15.7986 4 8.75029 8.93638 5.66406 16"/><path fill="#2F88FF" d="M24 14C20.2973 17.6298 18 22.6881 18 28.2829C18 28.7833 18.0184 29.2795 18.0545 29.7708C21.7253 33.3967 24 38.4327 24 44C24 38.4327 26.2747 33.3967 29.9455 29.7708C29.9816 29.2795 30 28.7833 30 28.2829C30 22.6881 27.7027 17.6298 24 14Z"/><path fill="#2F88FF" d="M4 24C4 35.0457 12.9543 44 24 44C24 38.4327 21.7253 33.3967 18.0545 29.7708C14.4424 26.2027 9.47841 24 4 24Z"/><path fill="#2F88FF" d="M44 24C44 35.0457 35.0457 44 24 44C24 38.4327 26.2747 33.3967 29.9455 29.7708C33.5576 26.2027 38.5216 24 44 24Z"/></g>`),
		g.Group(children),
	)
}