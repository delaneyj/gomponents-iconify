package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPakistan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#186648" d="M5 17h62v38H5z"/><path fill="#fff" d="M5 17h21v38H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M50 44a11.992 11.992 0 0 1-9.847-18.847a11.995 11.995 0 1 0 16.694 16.694A11.936 11.936 0 0 1 50 44Zm3.453-14.48l1.347-1.854l-2.179.708l-1.348-1.854v2.292l-2.179.708l2.179.708v2.292l1.348-1.854l2.179.708l-1.347-1.854z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}