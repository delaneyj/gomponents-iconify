package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kiss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zm-4.5 7a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3zm9 0a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3zM15 17.008v1.49c.901 0 1.477.415 1.477.63c0 .216-.572.626-1.467.63H15v1.492h.01c.895.003 1.467.414 1.467.629c0 .216-.576.63-1.477.63V24c1.669 0 2.977-.932 2.977-2.121c0-.533-.274-1.006-.713-1.375c.44-.369.713-.843.713-1.375c0-1.19-1.308-2.121-2.977-2.121z"/>`),
		g.Group(children),
	)
}