package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Presentation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.368 639h-896q-26 0-45-18.5t-19-45.5V192q0-27 19-45.5t45-18.5h896q26 0 45 18.5t19 45.5v384q0 26-19 44.5t-45 18.5zm32-575h-960q-13 0-22.5-9.5t-9.5-23T9.868 9t22.5-9h960q13 0 22.5 9t9.5 22.5t-9.5 23t-22.5 9.5zm-960 639h960q13 0 22.5 9.5t9.5 23t-9.5 22.5t-22.5 9h-326l95 177q11 19 5.5 40.5t-24.5 32.5t-40 5t-32-25l-125-230h-66l-125 230q-11 19-32 25t-40-5t-24.5-32.5t5.5-40.5l95-177h-326q-13 0-22.5-9t-9.5-22.5t9.5-23t22.5-9.5z"/>`),
		g.Group(children),
	)
}