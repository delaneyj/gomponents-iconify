package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angularjs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#B3B3B3" d="M63.81 1.026L4.553 21.88l9.363 77.637l49.957 27.457l50.214-27.828l9.36-77.635z"/><path fill="#A6120D" d="M117.536 25.998L63.672 7.629v112.785l45.141-24.983z"/><path fill="#DD1B16" d="m11.201 26.329l8.026 69.434l44.444 24.651V7.627z"/><path fill="#F2F2F2" d="m78.499 67.67l-14.827 6.934H48.044l-7.347 18.374l-13.663.254l36.638-81.508L78.499 67.67zm-1.434-3.491L63.77 37.858L52.864 63.726h10.807l13.394.453z"/><path fill="#B3B3B3" d="m63.671 11.724l.098 26.134l12.375 25.888H63.698l-.027 10.841l17.209.017l8.042 18.63l13.074.242z"/>`),
		g.Group(children),
	)
}