package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HostSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityHostSolid0" fill="currentColor"><circle cx="18" cy="25.64" r="1.5"/><path d="M26.5 1.86h-17A1.5 1.5 0 0 0 8 3.36v30.5h20V3.36a1.5 1.5 0 0 0-1.5-1.5ZM18 28.64a3 3 0 1 1 3-3a3 3 0 0 1-3 3Zm6-17H12v-1.6h12Zm0-4H12V6.06h12Z"/></g>`),
		g.Group(children),
	)
}