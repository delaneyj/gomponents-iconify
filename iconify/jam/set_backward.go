package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.18 8.477a.976.976 0 0 0 .324 1.363l9.922 6.032c.16.097.346.149.535.149c.56 0 1.013-.443 1.013-.99V2.97a.973.973 0 0 0-.153-.523a1.027 1.027 0 0 0-1.395-.317L6.504 8.16c-.131.08-.243.189-.325.317zM6 6.136L15.355.449c1.425-.867 3.3-.44 4.186.951c.3.471.459 1.014.459 1.569V15.03c0 1.641-1.36 2.97-3.04 2.97a3.093 3.093 0 0 1-1.605-.448L6 11.865V16a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v4.136zM2 2v14h2V2H2z"/>`),
		g.Group(children),
	)
}