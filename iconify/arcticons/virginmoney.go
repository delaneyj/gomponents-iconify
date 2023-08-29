package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Virginmoney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="31" x="4.5" y="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.348 20.4H43.5v7.2h-7.152a2.848 2.848 0 0 1-2.848-2.848v-1.504a2.848 2.848 0 0 1 2.848-2.848z"/><circle cx="36.472" cy="24" r=".795" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 27.545s2.11-3.932 2.617-4.468c0 0 2.363 7.506 2.617 10.485c.59-3.336 4.052-17.038 7.766-19.123"/>`),
		g.Group(children),
	)
}