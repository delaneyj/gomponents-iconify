package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 1024H64q-26 0-45-18.5T0 960V512q0-27 19-45.5T64 448h64V320q0-87 43-160.5T287.5 43T448 0t160.5 43T725 159.5T768 320v128h64q27 0 45.5 18.5T896 512v448q0 27-18.5 45.5T832 1024zm-512-64h256l-81-201q36-15 58.5-47t22.5-72q0-53-37.5-90.5T448 512t-90.5 37.5T320 640q0 40 22.5 72t58.5 47zm320-640q0-80-56-136t-135.5-56t-136 56T256 320v128h384V320z"/>`),
		g.Group(children),
	)
}