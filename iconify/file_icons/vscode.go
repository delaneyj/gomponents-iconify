package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vscode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M383.998 439.738V0L512 53.247v405.504L384 512L0 382.707l383.998 57.031zM61.18 299.748L32 290.001l72.036-71.208l-72.034-71.207l29.178-9.744l70.81 53.321l117.602-116.25l70.406 29.929v227.895l-70.41 29.928l-117.602-116.243l-70.806 53.324zm107.502-80.956l80.91 60.928V157.864l-80.91 60.928z"/>`),
		g.Group(children),
	)
}