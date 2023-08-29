package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icomoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M127.711 34.453c-170.281 98.17-170.281 344.925 0 443.094S512 452.338 512 256S297.992-63.716 127.711 34.453zm61.987 445.188c-178.705-108.502-164.206-350.92-.003-447.282c161.146 94.623 183.184 335.3.003 447.282zM248.218 256c0 45.42-49.508 73.963-88.901 51.253s-39.393-79.795 0-102.506s88.901 5.832 88.901 51.253z"/>`),
		g.Group(children),
	)
}