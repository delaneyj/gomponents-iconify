package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reinstead(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 2.5l-3.503 17.997L24 24l3.503-3.503L24 2.5zM24 24V2.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m45.5 24l-17.997-3.503L24 24l3.503 3.503L45.5 24zM24 24h21.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 45.5l3.503-17.997L24 24l-3.503 3.503L24 45.5zM24 24v21.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5 24l17.997 3.503L24 24l-3.503-3.503L2.5 24zM24 24H2.5m38.983-.782a17.501 17.501 0 0 0-16.7-16.7m-.001 34.965a17.501 17.501 0 0 0 16.7-16.701m-34.965 0a17.501 17.501 0 0 0 16.7 16.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.218 6.517a17.501 17.501 0 0 0-16.7 16.701m20.985 4.285l5.629 5.629m-12.635-5.629l-5.629 5.629m5.629-12.635l-5.629-5.629m12.635 5.629l5.629-5.629"/>`),
		g.Group(children),
	)
}