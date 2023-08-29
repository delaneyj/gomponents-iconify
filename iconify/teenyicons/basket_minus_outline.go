package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m2.449 14.398l.447-.224l-.447.224Zm10.102 0l.447.224l-.447-.224ZM.703 6h13.594V5H.703v1ZM14 5.703v.439h1v-.439h-1ZM12.386 14H2.614v1h9.772v-1ZM1 6.142v-.439H0v.439h1Zm1.896 8.032A17.96 17.96 0 0 1 1 6.142H0c0 2.944.685 5.847 2.002 8.48l.894-.448ZM2.614 14c.12 0 .229.068.282.174l-.894.448a.685.685 0 0 0 .612.378v-1Zm9.49.174a.315.315 0 0 1 .282-.174v1c.26 0 .496-.146.612-.378l-.894-.448ZM14 6.142c0 2.788-.65 5.538-1.896 8.032l.894.448A18.96 18.96 0 0 0 15 6.142h-1ZM14.297 6A.297.297 0 0 1 14 5.703h1A.703.703 0 0 0 14.297 5v1ZM.703 5A.703.703 0 0 0 0 5.703h1A.297.297 0 0 1 .703 6V5Zm3.226.757l3-5L6.07.243l-3 5l.858.514Zm4.142-5l3 5l.858-.514l-3-5l-.858.514ZM5 10h5V9H5v1Z"/>`),
		g.Group(children),
	)
}