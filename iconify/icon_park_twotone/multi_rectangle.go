package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMultiRectangle0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M12 39h32V7H12v8"/><path d="M8 39h24V15H8v8"/><path fill="#555" d="M20 23H4v16h16V23Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMultiRectangle0)"/>`),
		g.Group(children),
	)
}