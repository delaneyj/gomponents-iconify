package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Escalator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 320h-83L368 765l-7.5 7.5l-8.5 8.5l-9 8.5l-10.5 9l-11 8l-12 8.5l-12 7l-13.5 5.5l-13.5 3L256 832H128q-53 0-90.5-37.5T0 704t37.5-90.5T128 576h81l175-175V256q0-27 19-45.5t45.5-18.5t45 18.5T512 256v17l147-147q17-17 29-27.5T723.5 76T768 64h128q53 0 90.5 37.5T1024 192t-37.5 90.5T896 320zm0-192H768q-26 0-50 18L224 640h-96q-27 0-45.5 18.5t-18.5 45T82.5 749t45.5 19h128q25 0 49-20l493-492h98q27 0 45.5-19t18.5-45.5t-18.5-45T896 128zm-447.5 0q-26.5 0-45.5-19t-19-45.5t19-45T448.5 0t45 18.5t18.5 45t-18.5 45.5t-45 19z"/>`),
		g.Group(children),
	)
}