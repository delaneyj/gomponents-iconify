package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarineWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 20H7a1 1 0 0 0-.97 1.242L6.97 25h2.06l-.75-3h15.44l-.75 3h2.06l.94-3.758A1 1 0 0 0 25 20H14V3a1 1 0 0 0-1.8-.601L2.209 15.39A1 1 0 0 0 3 17h9Zm-6.97-5L12 5.94V15Z"/><path fill="currentColor" d="M26.906 25.751A2.998 2.998 0 0 1 24 28a3.438 3.438 0 0 1-3.051-2.316A1.007 1.007 0 0 0 20 25a.99.99 0 0 0-.947.679A3.44 3.44 0 0 1 16 28a3.438 3.438 0 0 1-3.051-2.316A1.007 1.007 0 0 0 12 25a.971.971 0 0 0-.947.679A3.44 3.44 0 0 1 8 28a2.998 2.998 0 0 1-2.906-2.249l-1.936.498A4.997 4.997 0 0 0 8 30a4.932 4.932 0 0 0 4-1.987a5.02 5.02 0 0 0 8 0A4.932 4.932 0 0 0 24 30a4.997 4.997 0 0 0 4.842-3.751Z"/><path fill="none" d="M23.75 10h-1.5V6h1.5ZM23 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M29.912 13.935L23.628 2.371a.718.718 0 0 0-1.256 0l-6.284 11.564A.72.72 0 0 0 16.72 15h12.56a.72.72 0 0 0 .631-1.065ZM22.25 6h1.5v4h-1.5Zm.75 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}