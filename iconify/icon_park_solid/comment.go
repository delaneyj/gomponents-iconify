package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSComment0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M44 6H4v30h9v5l10-5h21V6Z"/><path stroke="#000" d="M14 19.5v3m10-3v3m10-3v3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSComment0)"/>`),
		g.Group(children),
	)
}