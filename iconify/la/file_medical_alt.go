package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMedicalAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v11h1.367L8 12.42V5h10v6h6v16H8v-5.754l-.053-.053L6.756 20H6v9h20V9.594l-.281-.313l-6-6L19.406 3H6zm14 3.438L22.563 9H20V6.437zm-9.031 3.949l-2.336 5.832L8.414 16H2v2h5.586l1.777 1.781l1.668-4.168l3 7l2.07-5.175l.282.562h1.887a2 2 0 1 0 0-2h-.653l-1.719-3.438l-1.93 4.825l-3-7z"/>`),
		g.Group(children),
	)
}