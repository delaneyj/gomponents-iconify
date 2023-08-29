package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Idea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 511q-63 67-95.5 121.5T448 736q0 13-9.5 22.5T416 768H224q-13 0-22.5-9.5T192 736q0-49-32.5-103.5T64 511Q0 426 0 320q0-87 43-160.5T159.5 43T320 0t160.5 43T597 159.5T640 320q0 106-64 191zM224 832h192q13 0 22.5 9.5T448 864t-9.5 22.5T416 896H224q-13 0-22.5-9.5T192 864t9.5-22.5T224 832zm96 192q-53 0-90.5-19T192 960h256q0 26-37.5 45t-90.5 19z"/>`),
		g.Group(children),
	)
}