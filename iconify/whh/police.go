package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Police(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.35 256q0 62 32 187t32 197q0 110-60.5 183t-163.5 73q-40 0-77.5 13.5t-63 32t-45 37.5t-29 31.5t-9.5 13.5q-3-5-9.5-13.5t-28.5-31t-46-39t-62-30.5t-78-14q-103 0-163.5-73T.35 640q0-72 32-197t32-187q0-67-14.5-112T.35 64q0-3 .5-7t4-15t9.5-19.5t19-15.5t31-7q0 16 11 37t30.5 42t51 35t67.5 14q100 0 146-35t46-93h64q0 58 46 93t146 35q36 0 67.5-14t51-35t30.5-42t11-37q18 0 31 6.5t19 16t9.5 19t4.5 15.5v7q-35 35-49.5 80t-14.5 112zm-307 172l-77-172l-77 172l-179 24l132 130l-34 186l158-92l158 92l-34-186l132-130z"/>`),
		g.Group(children),
	)
}