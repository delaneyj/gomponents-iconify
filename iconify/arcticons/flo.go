package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 37.407c12.938.41 17.915-4.918 21.427-9.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.45 36.002c5.854-.235 14.05-5.386 14.05-14.168S38.695 5.91 28.45 5.91c-7.904 0-10.539 7.494-10.304 15.456m.058 20.724c-6.264-.351-8.486-4.765-8.486-4.765c2.163-3.197 1.578-7.236.758-9.578m22.071 9.578c-3.044 3.022-12.586 4.355-15.572 1.604m-.761-20.959c-2.928 2.4-3.981 6.674-2.87 11.416"/>`),
		g.Group(children),
	)
}