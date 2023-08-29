package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpectrumFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.2 2.006c8.041.088 8.8 1.245 8.8 9.995l-.005 1.199C21.908 21.24 20.75 22 12 22l-1.2-.005c-7.658-.083-8.712-1.137-8.794-8.795L2 11.69l.006-.89c.085-7.85 1.19-8.76 9.381-8.799l1.812.004ZM8.25 7.001h-.583a.667.667 0 0 0-.66.568l-.006.098v3.667c0 .335.246.612.568.66l.098.007h.584a3.75 3.75 0 0 1 3.744 3.55l.006.2v.583c0 .335.246.612.568.66l.098.007h3.667a.667.667 0 0 0 .66-.569l.007-.098v-.583a8.75 8.75 0 0 0-8.492-8.747l-.258-.003Z"/>`),
		g.Group(children),
	)
}