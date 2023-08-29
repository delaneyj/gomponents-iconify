package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noisecapture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.977 18.205c-1.095 0-1.983.847-1.983 1.89s.888 1.89 1.983 1.89s1.984-.845 1.984-1.89s-.888-1.89-1.984-1.89h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.504 33.776c-.12.171-.336.197-.482.056l-.007-.007l-.043-.049C8.846 32.256 4.5 25.984 4.5 20.791c-.001-3.667 2.542-6.641 5.68-6.643c3.138 0 5.683 2.972 5.684 6.639v.004c0 5.127-4.283 11.391-5.36 12.985h0ZM36.994 9.95c2.311 0 4.184 2.302 4.184 5.141v6.364c0 2.84-1.872 5.141-4.183 5.141s-4.183-2.301-4.183-5.141V15.09c0-2.839 1.873-5.14 4.183-5.14Zm1.617 19.356v4.219m-3.232 0v-4.22"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.56 22.77c.712 4.391 4.185 7.242 7.759 6.367c2.615-.64 4.66-3.153 5.181-6.367m-6.47-12.526v5.877m1.493-5.623v5.623m-2.986-5.554v5.554M10.155 33.938l6.166-.136l2.33-7.4l2.056 11.374l1.645-8.085l1.918 6.715l2.33-11.1l2.192 12.744l1.37-4.385l8.433.066"/>`),
		g.Group(children),
	)
}