package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneCleanGel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.022.794h.665m-.665 0V3.75m4.464-1.382l.94.94m-.47-.47l-2.114 2.114M23 7.107v1.33m0-.665h-2.991m1.417 4.464l-.94.94m.47-.47l-2.114-2.114m-2.155 4.158h-1.33m.665 0v-2.991M10 20.25H4A108.39 108.39 0 0 1 1 1.536A.751.751 0 0 1 1.75.75h10.5a.752.752 0 0 1 .75.786a108.395 108.395 0 0 1-3 18.714v0Zm0 0v2.25a.75.75 0 0 1-.75.75h-4.5A.75.75 0 0 1 4 22.5v-2.25m12-16.5a4 4 0 1 1-1.221 7.81M7 4.75v6m-3-3h6"/>`),
		g.Group(children),
	)
}