package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amigo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.997" cy="23.05" r="5.167" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.516 8.524L6.242 34.577c-3.185 6.253 3.275 10.083 6.87 6.609l3.464-3.348c5.847-5.65 9.229-5.53 14.963.204l2.474 2.474c4.416 4.416 10.917.144 7.773-5.968L28.395 8.524c-2.074-4.031-6.824-4.033-8.879 0Z"/>`),
		g.Group(children),
	)
}