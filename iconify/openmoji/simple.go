package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="25" fill="#FCEA2B" transform="rotate(-45.001 36 36)"/><path fill="#F1B31C" d="M37.405 57.829a54.578 54.578 0 0 1-3.511 3.2c7.054.544 14.305-1.873 19.713-7.281c9.802-9.802 9.834-25.662.07-35.426c-.166-.166-.348-.307-.519-.469c.502 14.392-4.744 28.967-15.753 39.976z"/><circle cx="36" cy="36" r="25" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/>`),
		g.Group(children),
	)
}