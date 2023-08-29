package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreetsBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m21.012 4.048l-7.952 7.951l7.952 7.952C22 18.455 22 16.134 22 12c0-4.133 0-6.455-.988-7.951ZM12 13.06l-7.952 7.952C5.544 22 7.866 22 11.999 22c4.135 0 6.456 0 7.952-.988L12 13.06Z"/><path d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.134 0 6.455.988 7.952L19.952 2.988C18.455 2 16.134 2 12 2C7.286 2 4.929 2 3.464 3.464Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}