package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Antenna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m865.417 119l-353 467v374q0 26-18.5 45t-45.5 19t-45.5-19t-18.5-45V586l-353-467q-31-19-31-55q0-26 19-45t45-19h768q27 0 45.5 19t18.5 45q0 36-31 55zm-153 9h-200v265zm-328 0h-200l200 265V128z"/>`),
		g.Group(children),
	)
}