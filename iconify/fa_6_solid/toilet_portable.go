package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToiletPortable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 32v32h320V32c0-17.7-14.3-32-32-32H32C14.3 0 0 14.3 0 32zm24 64H0v392c0 13.3 10.7 24 24 24s24-10.7 24-24v-8h224v8c0 13.3 10.7 24 24 24s24-10.7 24-24V96H24zm232 144v64c0 8.8-7.2 16-16 16s-16-7.2-16-16v-64c0-8.8 7.2-16 16-16s16 7.2 16 16z"/>`),
		g.Group(children),
	)
}