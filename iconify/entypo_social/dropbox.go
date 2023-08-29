package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.109.902L.4 4.457l3.911 3.279L10 4.043L6.109.902zm7.343 15.09a.44.44 0 0 1-.285-.102L10 13.262l-3.167 2.629a.447.447 0 0 1-.529.03l-2.346-1.533v.904L10 19.098l6.042-3.807v-.904l-2.346 1.533a.44.44 0 0 1-.244.072zM19.6 4.457L13.89.902L10 4.043l5.688 3.693L19.6 4.457zM10 11.291l3.528 2.928l5.641-3.688l-3.481-2.795L10 11.291zm-3.528 2.928L10 11.291L4.311 7.736l-3.48 2.795l5.641 3.688z"/>`),
		g.Group(children),
	)
}