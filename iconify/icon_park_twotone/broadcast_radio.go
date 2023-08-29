package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BroadcastRadio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBroadcastRadio0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 12h40v28H4z"/><path fill="#555" d="M31 31a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" d="M12 22h6m-6 8h6M8 40v4m32-4v4M8 12l28-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBroadcastRadio0)"/>`),
		g.Group(children),
	)
}