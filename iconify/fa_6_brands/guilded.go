package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Guilded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M443.427 64H4.571c0 103.26 22.192 180.06 43.418 222.358C112.046 414.135 224 448 225.256 448a312.824 312.824 0 0 0 140.55-103.477c25.907-33.923 53.1-87.19 65.916-145.761H171.833c4.14 36.429 22.177 67.946 45.1 86.944h88.589c-17.012 28.213-48.186 54.4-80.456 69.482c-31.232-13.259-69.09-46.544-96.548-98.362c-26.726-53.833-27.092-105.883-27.092-105.883h336.147A625.91 625.91 0 0 0 443.427 64Z"/>`),
		g.Group(children),
	)
}