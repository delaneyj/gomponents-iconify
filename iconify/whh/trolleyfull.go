package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trolleyfull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.12 960h-609q-22 30-55.5 47t-71.5 17q-66 0-113-47t-47-113q0-58 36.5-102t91.5-55V192q0-26-18.5-45t-45.5-19h-64q-26 0-45-18.5t-19-45T18.62 19t45.5-19h64q80 0 136 56t56 136v545q49 37 61 95h579q26 0 45 18.5t19 45t-19 45.5t-45 19zm-96-192h-448q-13 0-22.5-9.5t-9.5-22.5V608q0-13 9.5-22.5t22.5-9.5h448q13 0 22.5 9.5t9.5 22.5v128q0 13-9.5 22.5t-22.5 9.5zm-128-256h-320q-13 0-22.5-9.5t-9.5-22.5V352q0-13 9.5-22.5t22.5-9.5h320q13 0 22.5 9.5t9.5 22.5v128q0 13-9.5 22.5t-22.5 9.5zm-128-256h-192q-13 0-22.5-9.5t-9.5-22.5V96q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v128q0 13-9.5 22.5t-22.5 9.5z"/>`),
		g.Group(children),
	)
}