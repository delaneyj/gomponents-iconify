package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungCameraAssistant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.5 28.5h15a4 4 0 0 0 4-4v-15a4 4 0 0 0-4-4h-15a4 4 0 0 0-4 4v15a4 4 0 0 0 4 4Zm-12.5-2V37m0-31.5v10"/><circle cx="11" cy="21" r="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.5" cy="37" r="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 37h-7.5m26 0H35"/><circle cx="31" cy="17" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37" cy="11" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}