package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CancelTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.087 0 11 4.913 11 11c0 2.726-.99 5.206-2.625 7.125L9.03 7.47A10.954 10.954 0 0 1 16 5zM7.625 8.875L22.97 24.53A10.954 10.954 0 0 1 16 27C9.913 27 5 22.087 5 16c0-2.726.99-5.206 2.625-7.125z"/>`),
		g.Group(children),
	)
}