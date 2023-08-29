package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M432 80V16h-32v64h-64v32h64v64h32v-64h64V80h-64zM112 256h40v40h-40zm136 0h40v40h-40zm-44.562 128h-6.875a63.691 63.691 0 0 1-59.326-40h-34.47l4.662 11.653A95.541 95.541 0 0 0 196.563 416h6.875a95.54 95.54 0 0 0 89.133-60.347L297.233 344h-34.47a63.691 63.691 0 0 1-59.325 40Z"/><path fill="currentColor" d="M200 128C98.542 128 16 210.542 16 312s82.542 184 184 184s184-82.542 184-184s-82.542-184-184-184Zm0 336c-83.813 0-152-68.187-152-152s68.187-152 152-152s152 68.187 152 152s-68.187 152-152 152Z"/>`),
		g.Group(children),
	)
}