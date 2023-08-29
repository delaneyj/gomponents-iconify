package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaiduMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.886 22.647c-2.906-4.643-.813-12.473 3.618-15.691a12.472 12.472 0 0 1 14.862 0c4.733 3.622 6.612 10.908 3.706 15.55L23.935 40.302ZM17.111 43.5h13.8"/><circle cx="24.012" cy="17.947" r="5.531" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}