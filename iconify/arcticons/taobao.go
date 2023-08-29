package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taobao(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 10.403a15.796 15.796 0 0 1 4.827 4.291m-4.559 4.881c10.529 7.165 6.8 13.964.805 20.704M19.412 7.721c-.267 8.067-3.715 10.308-6.544 13.034"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.643 12.882a48.907 48.907 0 0 1 13.266-2.64c8.201-.296 10.77 4.217 11.317 9.601c.834 8.192-.151 13.39-4.13 17.326c-2.316 2.29-6.314 4.882-14.053.482M20.88 20.12h12.658m-6.978.211v13.417m-10.206-7.253h19.47m-4.076 2.037l1.716 6.008"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.711 31.905c-2.233.704-13.314 5.632-15.767-2.353m6.597-12.98l-5.149 6.865"/>`),
		g.Group(children),
	)
}