package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mymusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.38 896h-768q-53 0-90.5-37.5T.38 768V256q0-26 19-45t45-19h480l46-84q7-19 28-31.5t44-12.5h240q24 0 45 12.5t29 31.5l48 84v576q0 53-37.5 90.5t-90.5 37.5zm-296-515q-1 0-11-7t-16-11t-18-11t-21-12t-19.5-10t-19-7.5t-15.5-2.5q-13 0-22.5 10t-9.5 24v235q-30-13-64-13q-53 0-90.5 28t-37.5 68t37.5 68t90.5 28t90.5-28t37.5-68V512q0-21 9-42.5t23-21.5q20 0 36 32t22 64l6 32q-23 0-23 34q0 11 9 20.5t21 9.5q25 0 43-7.5t27-17t13.5-29t5-31.5t.5-37q0-23-10.5-45.5t-30-40.5t-32-28t-31.5-23zm-554-337q7-19 28-31.5t44-12.5h240q24 0 45 12.5t29 31.5l48 84H.38z"/>`),
		g.Group(children),
	)
}