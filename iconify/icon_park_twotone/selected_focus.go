package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectedFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSelectedFocus0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M12 4H4v8h8V4Zm32 32h-8v8h8v-8Zm-32 0H4v8h8v-8ZM44 4h-8v8h8V4Z"/><path stroke-dasharray="1 5" stroke-linecap="round" d="M8 36V12m32 24V12M12 8h24M12 40h24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSelectedFocus0)"/>`),
		g.Group(children),
	)
}