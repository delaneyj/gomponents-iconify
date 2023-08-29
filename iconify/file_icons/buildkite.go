package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buildkite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 169.844L342.694 256V85.032L512 169.844zM330.694 90.948l-154.031 75.938V337.83l154.03-75.916V90.948zm-166.031 75.914L0 85.032V256l164.663 81.807V166.862zm178.03 102.602v157.504L512 340.788v-157.48l-169.306 86.156z"/>`),
		g.Group(children),
	)
}