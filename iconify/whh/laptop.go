package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 832H64q-26 0-45-18.5t-19-45T19 723t45-19h896q26 0 45 19t19 45.5t-19 45t-45 18.5zm-64-192H128q-26 0-45-18.5T64 576V64q0-26 19-45t45-19h768q27 0 45.5 19T960 64v512q0 27-18.5 45.5T896 640zm0-544q0-13-9.5-22.5T864 64H160q-13 0-22.5 9.5T128 96v384q0 13 9.5 22.5T160 512h704q13 0 22.5-9.5T896 480V96zm-704 31h384L192 448V127z"/>`),
		g.Group(children),
	)
}