package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraPartyMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 45q18 0 30.5 12.5T427 88v256q0 18-12.5 30.5T384 387H43q-18 0-30.5-12.5T0 344V88q0-18 12.5-30.5T43 45h67l39-42h128l39 42h68zm-171 64q-44 0-75 31.5T107 216q0 10 2 21h44q-4-10-4-21q0-27 19-45.5t45-18.5h85q-32-43-85-43zm0 214q44 0 75.5-31.5T320 216q0-12-2-21h-45q4 10 4 21q0 27-18.5 45.5T213 280h-85q33 43 85 43z"/>`),
		g.Group(children),
	)
}