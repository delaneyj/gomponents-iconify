package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="14" height="18" x="34.607" y="3.494" rx="2" transform="rotate(45 34.607 3.494)"/><rect width="14" height="18" x="16.223" y="21.879" rx="2" transform="rotate(45 16.223 21.879)"/><path stroke-linecap="round" d="M31.072 16.929L16.93 31.071"/></g>`),
		g.Group(children),
	)
}