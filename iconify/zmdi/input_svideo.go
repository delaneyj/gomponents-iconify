package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputSvideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 229.5q0 13.5-9 22.5t-22.5 9t-23-9t-9.5-22.5t9.5-23t23-9.5t22.5 9.5t9 23zm150-107q0 13.5-9.5 23T267 155h-64q-14 0-23-9.5t-9-23t9-22.5t23-9h64q13 0 22.5 9t9.5 22.5zM160 304q13 0 22.5 9.5T192 336t-9.5 22.5T160 368t-22.5-9.5T128 336t9.5-22.5T160 304zM235 5q97 0 165.5 69T469 240t-68.5 166T235 475T69 406T0 240T69 74T235 5zm-.5 427q79.5 0 136-56.5T427 240t-56.5-135.5t-136-56.5T99 104.5T43 240t56 135.5T234.5 432zM352 197q13 0 22.5 9.5t9.5 23t-9.5 22.5t-22.5 9t-22.5-9t-9.5-22.5t9.5-23T352 197zm-42.5 107q13.5 0 22.5 9.5t9 22.5t-9 22.5t-22.5 9.5t-23-9.5T277 336t9.5-22.5t23-9.5z"/>`),
		g.Group(children),
	)
}