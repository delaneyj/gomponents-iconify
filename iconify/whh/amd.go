package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1015.697 1015l-247-247V288q0-13-9.5-22.5t-22.5-9.5h-480L9.697 9q10-9 23-9h960q13 0 22.5 9.5t9.5 22.5v960q0 13-9 23zm-759-279q0 13 9.5 22.5t22.5 9.5h416l-256 256h-416q-13 0-22.5-9.5T.697 992V576l256-256v416z"/>`),
		g.Group(children),
	)
}