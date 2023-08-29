package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Videocrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 12.572v23.316l-8.627-9.127v8.601a1.935 1.935 0 0 1-1.93 1.94H6.439a1.935 1.935 0 0 1-1.939-1.93V12.637c0-1.074.865-1.94 1.939-1.94h26.496c1.074 0 1.939.866 1.939 1.94v9.06l8.626-9.126Z"/>`),
		g.Group(children),
	)
}