package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Financing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFinancing0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><rect width="10" height="10" x="24" y="16.929" rx="2" transform="rotate(45 24 16.929)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFinancing0)"/>`),
		g.Group(children),
	)
}