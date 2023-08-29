package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M338 973L75 677q-24-23-49.5-72T0 526V150Q0 88 44 44T149 0h598q62 0 105.5 44T896 150v376q0 31-25 82.5T822 677L558 973q-47 51-110 51t-110-51z"/>`),
		g.Group(children),
	)
}