package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.486.807l3.603 3.326l-.678.734L8 2.642V11H7V2.707L4.854 4.854l-.708-.708l3.34-3.34ZM2 13V7H1v7h13V7h-1v6H2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}