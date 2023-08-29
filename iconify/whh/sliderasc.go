package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sliderasc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M991.998 1024h-192q-13 0-22.5-9.5t-9.5-22.5V800q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm-32-448h-256V448h256q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19zm0-384h-640V64h640q27 0 45.5 18.5t18.5 45.5t-18.5 45.5t-45.5 18.5zm-736 64h-192q-13 0-22.5-9.5t-9.5-22.5V32q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm-160 192h256v128h-256q-26 0-45-19t-19-45.5t19-45t45-18.5zm352 192q-13 0-22.5-9.5t-9.5-22.5V416q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5h-192zm-352 192h640v128h-640q-26 0-45-19t-19-45.5t19-45t45-18.5z"/>`),
		g.Group(children),
	)
}