package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOfDavid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b399c8" d="M12 12.08h48v47.83H12z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="2" d="M12 12h48v48H12z"/><path stroke-width="2.696" d="m36 18.67l14.97 25.96H21.03l7.486-12.98z" paint-order="stroke fill markers"/><path stroke-width="2.696" d="m21.01 27.39l29.97.019l-14.97 25.93l-7.5-12.97z" paint-order="stroke fill markers"/></g>`),
		g.Group(children),
	)
}