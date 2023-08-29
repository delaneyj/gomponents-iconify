package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cpanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 512q0 27-19 45.5T960 576H576v384q0 27-18.5 45.5T512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0q139 0 257 68.5T955.5 255t68.5 257zm-896 0q0 142 91.5 248.5T448 890V134q-137 23-228.5 129.5T128 512zm448-378v314h314q-20-121-107-207.5T576 134z"/>`),
		g.Group(children),
	)
}