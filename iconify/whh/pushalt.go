package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pushalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512.59 512q-139 0-257-17T69.094 448.5T.59 384V256q0-8 6-19q62 37 202.5 60t303.5 23t303.5-23t202.5-60q6 11 6 19v128q0 35-68.5 64.5T769.59 495t-257 17zm0-256q-139 0-257-17T69.094 192.5T.59 128t68.5-64.5T255.59 17t257-17t257 17t186.5 46.5t68.5 64.5t-68.5 64.5t-186.5 46.5t-257 17zm-48 335q20.002-16 47.502-16t47.5 16l189 150q19 16 19 38t-19.5 37.5t-47 15.5t-47.5-16l-77-61v205q0 26-18.5 45t-45 19t-45.5-19t-19-45V754l-78 62q-20 16-47.5 16t-47-15.5t-19.5-37.5t19-38z"/>`),
		g.Group(children),
	)
}