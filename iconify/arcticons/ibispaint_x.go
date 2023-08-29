package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbispaintX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.875 33.5a14.5 14.5 0 1 1 14.5-14.5a14.517 14.517 0 0 1-14.5 14.5Zm0-21a6.5 6.5 0 1 0 6.5 6.5a6.508 6.508 0 0 0-6.5-6.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.375 31.956V39.5a4 4 0 0 1-8 0V19m-7 23.5V22"/><circle cx="6.375" cy="19" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.875 25.5v8m-4.596-9.904l-5.657 5.657M21.375 19h-8m9.904-4.596l-5.657-5.657M27.875 12.5v-8m4.596 9.904l5.657-5.657M34.375 19h8m-9.904 4.596l5.657 5.657"/>`),
		g.Group(children),
	)
}