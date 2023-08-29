package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scribd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M269 412q-2 2-5.5 6t-5.5 6q-39 38-96 38q-55 0-96-29q-42-30-65-89l1-1q-2-8-2-13q0-21 15-35.5T50 280q21 0 35.5 14.5T100 330q0 19-15 36q31 34 75 34q32 0 54-19q22-20 22-49q0-18-12-36q-11-18-28-29q-18-11-53-26q-38-16-56-27q-22-15-34-30q-14-16-21-34q-7-15-7-36q0-47 37-80q40-32 92-32q34 0 72 15q28 11 52 37q17 15 17 37q0 20-14.5 34.5T245 140t-35.5-14.5T195 91q0-9 5-21q-15-7-41-7q-32 0-52 14q-20 15-20 37q0 20 17 35q18 15 60 33q43 19 65 33q20 13 38 34q16 21 23 40q0 2 1 2q-35 29-35 73q0 25 13 48zm73-105q-24 0-41 17t-17 40q0 24 17 41t41 17t41-17t17-41q0-23-17-40t-41-17z"/>`),
		g.Group(children),
	)
}