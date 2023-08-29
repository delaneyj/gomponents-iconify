package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slidersoff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M959.998 576h-640V448h640q26 0 45 19t19 45t-19 45t-45 19zm0-384h-640V64h640q26 0 45 19t19 45.5t-19 45t-45 18.5zm-736 832h-192q-13 0-22.5-9.5t-9.5-22.5V800q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm0-384h-192q-13 0-22.5-9.5t-9.5-22.5V416q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm0-384h-192q-13 0-22.5-9.5t-9.5-22.5V32q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm800 640q0 26-19 45t-45 19h-640V832h640q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}