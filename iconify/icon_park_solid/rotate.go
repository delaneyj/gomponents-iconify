package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rotate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRotate0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M12 24h30v18H12V24Z"/><path stroke-linecap="round" d="M6 8v9h9"/><path stroke-linecap="round" d="M38.475 13.299C35.195 8.87 29.933 6 24 6c-5.821 0-10.997 2.763-14.287 7.05L6 17"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRotate0)"/>`),
		g.Group(children),
	)
}