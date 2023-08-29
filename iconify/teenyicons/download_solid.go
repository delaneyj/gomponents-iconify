package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 9.358V1h1v8.293l2.146-2.147l.708.708l-3.34 3.34L3.91 7.866l.678-.734L7 9.358ZM2 13V7H1v7h13V7h-1v6H2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}