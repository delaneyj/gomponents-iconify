package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mojo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="7.611" height="18.33" x="7.177" y="14.835" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.806" ry="3.806" transform="rotate(30 10.982 24)"/><rect width="7.611" height="18.33" x="20.194" y="14.835" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.806" ry="3.806" transform="rotate(30 24 24)"/><rect width="7.611" height="18.33" x="33.212" y="14.835" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.806" ry="3.806" transform="rotate(30 37.018 24)"/>`),
		g.Group(children),
	)
}