package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squeezer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.067 17.808h13.688m3.106 5.81v9.963a1.853 1.853 0 0 1-1.639 1.66l-1.875-.009l-.002 5.71a2.061 2.061 0 0 1-4.123.009l.002-5.71H28.1l-.002 5.71a2.061 2.061 0 0 1-4.123.008l.003-5.71l-1.875.007a1.884 1.884 0 0 1-1.542-1.185a1.294 1.294 0 0 1-.098-.475"/><circle cx="25.564" cy="12.273" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.67 8.03l1.367-3.044m-11.93 0l1.366 3.044"/><circle cx="20.041" cy="24.753" r="2.784" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="20.041" cy="30.631" r="2.784" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="20.041" cy="18.772" r="2.784" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.891" cy="19.801" r="3.609" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.517 15.869a9.7 9.7 0 0 1 19.308.168m-26.372 3.691a7.834 7.834 0 0 0-.027 10.01h.027M8.022 14.95a15.404 15.404 0 0 0 0 19.605"/><circle cx="34.671" cy="12.273" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}