package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MintMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.081 23.475c7.905-2.108 15.81-1.054 15.81 1.054h4.217c0-2.108 7.905-3.426 15.81-1.054"/><path d="M7.135 23.211c0 8.96 3.69 10.804 7.379 10.804s7.378-2.635 7.378-9.486m18.973-1.318c0 8.96-3.69 10.804-7.379 10.804s-7.378-2.635-7.378-9.486"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.69 15.358c4.215-2.635 8.959-6.324 14.756-5.533c.527 4.216-.527 7.905-2.108 12.121c1.054 2.372 2.371 4.48 3.162 8.96l-2.372-.264l2.372 3.426c-1.921-1.112-3.61-1.779-5.126-2.103c-3.219-.688-5.656.167-7.882 1.567c-2.187 1.375-4.17 3.275-6.492 4.752m-3.69-22.926c-2.634-1.58-5.27-3.689-7.905-4.48c-2.635-.526-6.851 4.48-6.545 5.214c.399 1.907 1.04 3.82 1.802 5.854c-1.054 2.372-2.371 4.48-3.162 8.96l2.372-.264L4.5 34.068c5.913-3.423 9.62-2.62 12.858-.629c2.247 1.381 4.267 3.334 6.641 4.845M20.31 15.358c2.636-.527 4.744-.527 7.38 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.784 35.122c2.635.264 5.797.264 8.432 0h0M5.86 16.092c2.066 1.111 4.437 1.638 6.545 1.638c.527-3.162.527-5.006 0-6.85"/>`),
		g.Group(children),
	)
}