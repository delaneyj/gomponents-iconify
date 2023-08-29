package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bubbleupnp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.25 41.457a25.17 25.17 0 0 0-1.654-34.914"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.216 36.83a18.304 18.304 0 0 0-1.19-25.142m-3.91 20.454a11.36 11.36 0 0 0-.73-15.42M15.25 24.36a5.751 5.751 0 1 1 0 11.497H5.5V12.86h9.75a5.848 5.848 0 0 1 5.91 5.749a5.848 5.848 0 0 1-5.91 5.749Zm0-.001H5.5"/>`),
		g.Group(children),
	)
}