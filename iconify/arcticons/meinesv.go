package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meinesv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.003 4.504V43.5m0 0a46.003 46.003 0 0 1-5.826-2.726a18.764 18.764 0 0 1-5.212-4.36a12.554 12.554 0 0 1-2.855-7.03c-.237-2.918.427-6.481 2.892-10.736s6.733-9.199 11.001-14.143"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.003 4.5c3.89 4.886 7.779 9.77 10.177 13.366s3.305 5.904 3.628 8.347a11.578 11.578 0 0 1-1.155 7.187a13.017 13.017 0 0 1-5.658 5.01a27.122 27.122 0 0 1-6.992 1.99"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.514 23.054l-4.129 10.941l-4.129-10.941M14.087 33.08c.755.635 1.57.925 3.4.925h.927a2.735 2.735 0 0 0 2.732-2.738h0a2.735 2.735 0 0 0-2.733-2.738H16.56a2.735 2.735 0 0 1-2.732-2.739h0a2.735 2.735 0 0 1 2.732-2.738h.927c1.83 0 2.645.29 3.4.924"/>`),
		g.Group(children),
	)
}