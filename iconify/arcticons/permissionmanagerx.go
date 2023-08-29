package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Permissionmanagerx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.337 28.576l-2.319 4.005a.749.749 0 0 1-.933.327l-3.817-1.532q-.23.176-.47.34l-4.079-4.23a5.713 5.713 0 1 0-8.04-4.186l-5.825-1.68q.027-.4.079-.798L5.707 18.3a.742.742 0 0 1-.187-.974l3.062-5.29a.749.749 0 0 1 .934-.327l3.83 1.533a11.487 11.487 0 0 1 2.585-1.488l.575-4.039a.746.746 0 0 1 .747-.64h6.14a.746.746 0 0 1 .746.64l.575 4.04a11.723 11.723 0 0 1 2.584 1.487l3.817-1.532a.749.749 0 0 1 .933.327l.706 1.22M21.807 30.868a5.713 5.713 0 0 1-8.04-4.185h0l-5.825-1.681q-.023.344-.026.689a12.235 12.235 0 0 0 .104 1.488L4.794 29.7a.742.742 0 0 0-.187.974l3.062 5.29a.748.748 0 0 0 .933.326l3.817-1.532a11.725 11.725 0 0 0 2.584 1.488l.575 4.039a.746.746 0 0 0 .747.64h6.124a.746.746 0 0 0 .747-.64l.575-4.04a11.484 11.484 0 0 0 2.114-1.147Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.41 28.498H43.5l-9.045-18.09Z"/><circle cx="34.455" cy="25.642" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.455 22.785V16.12"/>`),
		g.Group(children),
	)
}