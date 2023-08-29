package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphabeticalSortingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v2h5.563L5.28 12.28l-.28.314V15h8v-2H7.437l5.282-5.28l.28-.314V5H5zm17 0v18.688l-2.594-2.594L18 22.5l4.28 4.313l.72.687l.72-.688L28 22.5l-1.406-1.406L24 23.687V5h-2zM8.187 17l-.218.656L6.03 23H6v.063l-.938 2.593l-.062.157V27h2v-.844L7.406 25h3.188L11 26.156V27h2v-1.188l-.063-.156L12 23.063V23h-.03l-1.94-5.344L9.814 17H8.185zM9 20.656L9.844 23H8.156L9 20.656z"/>`),
		g.Group(children),
	)
}