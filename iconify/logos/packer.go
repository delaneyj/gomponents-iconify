package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Packer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#1DAEFF" d="m.357 50.596l135.137 77.266v284.835L.357 335.447"/><path fill="#1D94DD" d="M202.395 83.892L55.505.188v58.127l99.887 57.065v174.429l47.01 26.71c29.092 16.6 52.908 6.655 52.908-22.11V166.417c0-28.78-23.816-65.916-52.907-82.525"/>`),
		g.Group(children),
	)
}