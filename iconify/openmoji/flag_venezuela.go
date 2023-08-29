package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagVenezuela(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#f1b31c" d="M5 17h62v13H5z"/><path fill="#1e50a0" d="M5 30h62v12H5z"/><g fill="#fff"><circle cx="34" cy="33" r="1"/><circle cx="38" cy="33" r="1"/><circle cx="30" cy="34" r="1"/><circle cx="27" cy="36" r="1"/><circle cx="25" cy="39" r="1"/><circle cx="42" cy="34" r="1"/><circle cx="45" cy="36" r="1"/><circle cx="47" cy="39" r="1"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}