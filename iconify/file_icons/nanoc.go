package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nanoc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128.676 356.921V153.88h90.33l81.536 113.111V153.88h78.339v203.04h-89.13l-83.008-111.835l-1.326 111.835h-76.741zm-.965-322.468c-170.281 98.17-170.281 344.925 0 443.094S512 452.338 512 256S297.993-63.716 127.711 34.453zm26.439 397.435c-135.188-77.937-135.188-273.839 0-351.776C289.338 2.174 459.241 100.125 459.241 256S289.338 509.826 154.15 431.888z"/>`),
		g.Group(children),
	)
}