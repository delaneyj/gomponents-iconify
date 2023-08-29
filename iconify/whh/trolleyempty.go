package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trolleyempty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.43 960h-609q-22 30-55.5 47t-71.5 17q-66 0-113-47t-47-113q0-58 36.5-102t91.5-55V192q0-27-18.5-45.5t-45.5-18.5h-64q-26 0-45-19t-19-45.5t19-45t45-18.5h64q80 0 136 56t56 136v545q49 37 61 95h579q26 0 45 18.5t19 45t-18.5 45.5t-45.5 19z"/>`),
		g.Group(children),
	)
}