package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analogright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1011 469L829 599q-13 9-30.5 9t-30.5-9V297q13-9 31-9t30 9l182 129q13 9 13 21.5t-13 21.5zM704 239v418q0 20 17 31.5t38 11t34-10.5l50-31q-57 108-163 173t-232 65q-91 0-174-35.5T131 765T35.5 622T0 448t35.5-174T131 131t143-95.5T448 0q126 0 231.5 64.5T843 237l-50-30q-12-10-33-10.5T721.5 207T704 239z"/>`),
		g.Group(children),
	)
}