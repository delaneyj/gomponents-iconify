package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreelancerUpwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.043 22.04v6.91M12 18.354v6.651a4.044 4.044 0 0 0 4.033 4.033h0a4.042 4.042 0 0 0 4.03-4.033v-6.65m2.569-.016s2.222 10.655 8.54 10.474c2.51-.07 4.866-1.711 4.827-5.405c-.032-3.435-2.318-5.034-5.02-4.766c-7.268.724-6.748 19.496-6.748 19.496"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}