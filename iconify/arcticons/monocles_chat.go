package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonoclesChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.004 27.816s1.342-1.052 2.164-.65c2.756 1.347 3.96 5.414 9.71 2.302c-3.924 7.14-10.172 4.097-11.874 1.598c-1.701 2.5-7.95 5.541-11.873-1.598c5.749 3.112 6.954-.955 9.71-2.303c.821-.401 2.163.651 2.163.651Z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.423 27.343v-4.941"/><circle cx="30.826" cy="19.381" r="3.742" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.423" cy="19.381" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.578 27.343v-4.941"/><circle cx="17.174" cy="19.381" r="3.742" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="10.578" cy="19.381" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.15 34.36L2.5 45.5l11.14-2.65"/>`),
		g.Group(children),
	)
}