package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phoneprofilesplus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 32.545a.75.75 0 1 0 .75.75a.74.74 0 0 0-.73-.75H16Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.16 33.395a2 2 0 0 1-2 2h-8.32a2 2 0 0 1-2-2v-18.79a2 2 0 0 1 2-2h8.32a2 2 0 0 1 2 2v18.79ZM9.84 16.225h12.32m3.68-.5h12.32m-12.32 5.5h12.32m-12.32 5.5h12.32m-12.32 5.5h12.32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.84 30.685h12.32M18.109 18.75l-4.218 5.2h4.218l-3.912 4.21"/>`),
		g.Group(children),
	)
}