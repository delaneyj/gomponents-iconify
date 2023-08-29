package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSwitchButton0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M36 4H12a8 8 0 1 0 0 16h24a8 8 0 1 0 0-16Zm0 24H12a8 8 0 1 0 0 16h24a8 8 0 1 0 0-16Z"/><path d="M36 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM12 38a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSwitchButton0)"/>`),
		g.Group(children),
	)
}