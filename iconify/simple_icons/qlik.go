package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qlik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.752 20.137l-3.13-2.633a11.195 11.195 0 0 0 1.73-5.965C22.352 5.37 17.344.362 11.176.362S0 5.37 0 11.54c0 6.168 5.008 11.176 11.176 11.176c2.393 0 4.622-.756 6.444-2.044l3.333 2.799s.497.423.92-.074l1.989-2.357c-.019 0 .386-.497-.11-.902zm-5.708-8.598a6.867 6.867 0 1 1-13.735 0a6.867 6.867 0 0 1 13.735 0zm-11.287 0a4.4 4.4 0 1 1 8.8 0a4.4 4.4 0 0 1-8.8 0z"/>`),
		g.Group(children),
	)
}