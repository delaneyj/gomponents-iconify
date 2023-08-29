package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.295 8.362H7.705L4.5 11.568v7.915l3.205 3.205h32.59l3.205-3.205v-7.915l-3.205-3.206zm-4.674 2.78v8.766m-23.242-8.766v8.766m10.061-8.766h5.808m-2.904 8.766v-8.766m5.147 0h2.904m-2.904 8.766v-8.766"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.612 19.908v-8.766l5.807 8.766v-8.766"/><rect width="37.186" height="7.865" x="5.407" y="31.773" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.096"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.605 28.946h18.79"/>`),
		g.Group(children),
	)
}