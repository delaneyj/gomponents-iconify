package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pilcrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 128h-64v832q0 26-19 45t-45.5 19t-45-19t-18.5-45V128H512v832q0 26-19 45t-45.5 19t-45-19t-18.5-45V640h-64q-87 0-160.5-43T43 480.5T0 320t43-160.5T159.5 43T320 0h640q26 0 45 18.5t19 45t-18.5 45.5t-45.5 19z"/>`),
		g.Group(children),
	)
}