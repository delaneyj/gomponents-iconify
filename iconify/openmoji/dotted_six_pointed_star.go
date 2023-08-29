package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DottedSixPointedStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b399c8" d="M12 12.08h48v47.83H12z"/><path d="M38.5 36a2.5 2.5 0 1 1-2.499-2.5A2.499 2.499 0 0 1 38.5 36z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12h48v48H12z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.696" d="m36 18.67l14.97 25.96H21.03l7.486-12.98z" paint-order="stroke fill markers"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.696" d="m21.01 27.39l29.97.019l-14.97 25.93l-7.5-12.97z" paint-order="stroke fill markers"/>`),
		g.Group(children),
	)
}