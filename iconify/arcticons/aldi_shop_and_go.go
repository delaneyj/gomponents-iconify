package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AldiShopAndGo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.632 34.856c0-1.007-.873-1.88-1.88-1.812c-.94.067-1.677.94-1.677 1.88V36.6c0 1.007.805 1.812 1.812 1.812h0a1.805 1.805 0 0 0 1.812-1.812h-1.812m-13.473 1.21c.336.402.739.603 1.343.603h.805c.738 0 1.343-.604 1.343-1.342h0c0-.738-.604-1.342-1.343-1.342h-.872a1.346 1.346 0 0 1-1.343-1.343h0c0-.738.604-1.342 1.343-1.342h.805c.604 0 1.007.134 1.342.604m6.445 4.765h-.403c-.268 0-.604-.134-.738-.402l-1.88-2.55c-.268-.337-.67-.74-.67-1.343s.47-1.074 1.14-1.074c.605 0 1.074.47 1.074 1.074s-.47 1.14-1.41 1.342c-1.006.201-1.677.738-1.677 1.678c0 .738.47 1.275 1.41 1.275c1.208 0 1.946-1.208 3.02-2.617m-5.273-25.297L14.5 27.501m10.297-17.002l-6.375 17.002m10.297-17.002l-6.375 17.002m7.117-12.752h.871M27.867 19h4.039m-5.632 4.251H33.5"/>`),
		g.Group(children),
	)
}