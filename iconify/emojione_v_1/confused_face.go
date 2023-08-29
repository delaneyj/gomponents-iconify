package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConfusedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#fbbf67" d="M63.29 31.941c0 17.476-14.167 31.645-31.645 31.645c-17.48 0-31.649-14.169-31.649-31.645C-.004 14.461 14.165.292 31.645.292c17.478 0 31.645 14.169 31.645 31.649"/><ellipse cx="42.16" cy="27.514" fill="#fff" rx="6.432" ry="9.631"/><path fill="#25333a" d="M37.998 27.514c0 2.787 1.86 5.05 4.16 5.05c2.301 0 4.161-2.259 4.161-5.05c0-2.787-1.86-5.05-4.161-5.05c-2.3 0-4.16 2.26-4.16 5.05"/><ellipse cx="20.649" cy="27.19" fill="#fff" rx="6.432" ry="9.631"/><ellipse cx="20.649" cy="27.19" fill="#25333a" rx="4.161" ry="5.05"/><path fill="#633d19" d="M45.04 42.849c-10.166-1.612-21.15 1.373-30.25 5.839c-2.02.993-.592 4.192 1.442 3.195c8.56-4.201 18.866-7.07 28.44-5.55c2.234.355 2.591-3.131.366-3.484"/>`),
		g.Group(children),
	)
}