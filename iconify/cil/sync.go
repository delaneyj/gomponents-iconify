package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m410.168 133.046l-28.958-28.958l82.807-.088l-.034-32L328 72.144V208h32v-79.868l27.541 27.541A152.5 152.5 0 0 1 279.972 416l.056 32a184.5 184.5 0 0 0 130.14-314.954ZM232.028 104l-.056-32a184.5 184.5 0 0 0-130.14 314.954L130.878 416H48v32h136V312h-32v79.868l-27.541-27.541A152.5 152.5 0 0 1 232.028 104Z"/>`),
		g.Group(children),
	)
}