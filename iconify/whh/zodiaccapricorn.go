package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zodiaccapricorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800.43 960q-92 0-160-65q-94 129-256 129q-26 0-45-19t-19-45.5t19-45t45-18.5q21 0 32.5-.5t31-4t33-11.5t31-21.5t33-36.5t31.5-54l-55-611q-58 62-97.5 224t-39.5 323q0 26-18.5 45t-45 19t-45.5-19t-19-45q0-129-25.5-261.5T160.43 219t-96-91q-26 0-45-19t-19-45.5t19-45t45-18.5q50 0 99.5 48t87.5 122t69 169q42-130 99.5-221T541.43 6q12-6 25-6q1 0 2.5.5t2 0t2.5-.5h3v1q50 5 64 63l45 492q56-44 115-44q93 0 158.5 65.5t65.5 158.5t-65.5 158.5t-158.5 65.5zm0-320q-37 0-66.5 41t-29.5 74q0 30 28.5 53.5t67.5 23.5q40 0 68-28t28-68t-28-68t-68-28z"/>`),
		g.Group(children),
	)
}