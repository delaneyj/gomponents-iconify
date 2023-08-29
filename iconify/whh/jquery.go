package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jquery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M807 480q-109 0-186-77t-77-186q0-64 36-121.5T672 0q-4 10-12.5 31t-13 32t-12 29.5t-11 30T616 148t-6 24.5t-2 19.5q0 95 64.5 159.5T832 416q9 0 19.5-2t24.5-6t25.5-7.5t30-11t29.5-12t32-13t31-12.5q-38 56-95.5 92T807 480zm-71 160q61 0 114.5-7T992 608q-55 62-133 95t-167 33q-110 0-203-54T342 535t-54-202q0-90 33-168t95-133q-18 88-25 141.5T384 288q0 94 48 175t129 129t175 48zM576 896q183 0 320-32q-70 78-170 119t-214 41q-139 0-257-68.5T68.5 769T0 512q0-114 41-214t119-170q-32 137-32 320q0 91 35.5 174T259 765t143 95.5T576 896z"/>`),
		g.Group(children),
	)
}