package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slidersfull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992.232 1024h-192q-13 0-22.5-9.5t-9.5-22.5V800q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm0-384h-192q-13 0-22.5-9.5t-9.5-22.5V416q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm0-384h-192q-13 0-22.5-9.5t-9.5-22.5V32q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm-992 640q0-27 18.5-45.5t45.5-18.5h640v128h-640q-27 0-45.5-18.5T.232 896zm0-384q0-27 18.5-45.5t45.5-18.5h640v128h-640q-27 0-45.5-18.5T.232 512zm0-384q0-27 18.5-45.5t45.5-18.5h640v128h-640q-27 0-45.5-18.5T.232 128z"/>`),
		g.Group(children),
	)
}