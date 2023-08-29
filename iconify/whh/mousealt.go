package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mousealt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M320 1024q-87 0-160.5-43T43 864.5T0 704V320q0-87 43-160.5T159.5 43T320 0t160.5 43T597 159.5T640 320v384q0 87-43 160.5T480.5 981T320 1024zm.5-832q-26.5 0-45.5 19t-19 45t19 45t45.5 19t45-19t18.5-45t-18.5-45t-45-19z"/>`),
		g.Group(children),
	)
}