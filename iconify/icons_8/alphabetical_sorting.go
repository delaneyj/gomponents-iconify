package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphabeticalSorting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m8.188 5l-.22.656L6.032 11H6v.063l-.938 2.593l-.062.156V15h2v-.844L7.406 13h3.188L11 14.156V15h2v-1.188l-.063-.156L12 11.062V11h-.03l-1.94-5.344L9.814 5H8.185zM22 5v18.688l-2.594-2.594L18 22.5l4.28 4.313l.72.687l.72-.688L28 22.5l-1.406-1.406L24 23.687V5h-2zM9 8.656L9.844 11H8.156L9 8.656zM5 17v2h5.563L5.28 24.28l-.28.314V27h8v-2H7.437l5.282-5.28l.28-.314V17H5z"/>`),
		g.Group(children),
	)
}