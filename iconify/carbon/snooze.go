package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snooze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 14v-2h-8v2h5.5L12 20v2h8v-2h-5.493L20 14zm3.001-8.588L24.416 4L28 7.59l-1.415 1.412z"/><path fill="currentColor" d="M16 28a11 11 0 1 1 11-11a11.012 11.012 0 0 1-11 11Zm0-20a9 9 0 1 0 9 9a9.01 9.01 0 0 0-9-9ZM4.002 7.589l3.583-3.59L9 5.41L5.417 9z"/>`),
		g.Group(children),
	)
}