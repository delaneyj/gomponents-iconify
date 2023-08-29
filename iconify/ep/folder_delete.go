package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M128 192v640h768V320H485.76L357.504 192H128zm-32-64h287.872l128.384 128H928a32 32 0 0 1 32 32v576a32 32 0 0 1-32 32H96a32 32 0 0 1-32-32V160a32 32 0 0 1 32-32zm370.752 448l-90.496-90.496l45.248-45.248L512 530.752l90.496-90.496l45.248 45.248L557.248 576l90.496 90.496l-45.248 45.248L512 621.248l-90.496 90.496l-45.248-45.248L466.752 576z"/>`),
		g.Group(children),
	)
}