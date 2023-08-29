package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m201.37 288l176-176H48v352h352V134.63l-176 176L201.37 288z"/><path fill="currentColor" d="M320 48v32h89.37l-32 32L400 134.63l32-32V192h32V48H320z"/>`),
		g.Group(children),
	)
}