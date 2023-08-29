package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apache(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M991.998 1024q-25 0-52-1q-414-11-637.5-176t-236.5-460q-66-171-66-355q0-13 9.5-22.5t22.5-9.5t22.5 9.5t9.5 22.5v6q-9 15-4 32q2 15 8 40q4 42 11 84q2 0 5 1t4 1q78 322 168 412q60 61 110 106t112.5 91.5t127.5 78t130 44.5h64q-11-6-29.5-16.5t-75-47t-111.5-78t-128.5-107t-135.5-135.5q-79-89-151-307q45 34 87 83q92 108 246 257.5t299.5 266t190.5 116.5q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5z"/>`),
		g.Group(children),
	)
}