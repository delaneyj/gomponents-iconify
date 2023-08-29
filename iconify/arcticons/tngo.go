package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tngo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M20.778 20.92a8.723 8.723 0 1 1-3.193-3.192"/><path d="M14.25 23.505a2.052 2.052 0 1 1-2.053 0l1.027 1.778l1.026-1.778Zm21.553-.172a2.052 2.052 0 1 1-2.052 0l1.026 1.777l1.026-1.777Z"/><circle cx="34.777" cy="25.282" r="8.723"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.667 19.84l5.774-3.334m7.001-2.256h6.667m-28.221-.256h6.668"/>`),
		g.Group(children),
	)
}