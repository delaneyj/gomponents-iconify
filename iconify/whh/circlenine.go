package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlenine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm192-640q0-53-37.5-90.5T576 256H448q-53 0-90.5 37.5T320 384v64q0 53 37.5 90.5T448 576h128q34 0 64-17v81q0 26-18.5 45T576 704H448q-26 0-45-19t-19-45q0-13-9.5-22.5T352 608t-22.5 9.5T320 640q0 53 37.5 90.5T448 768h128q53 0 90.5-37.5T704 640V384zM576 512H448q-26 0-45-19t-19-45v-64q0-27 19-45.5t45-18.5h128q27 0 45.5 18.5T640 384v64q0 26-18.5 45T576 512z"/>`),
		g.Group(children),
	)
}