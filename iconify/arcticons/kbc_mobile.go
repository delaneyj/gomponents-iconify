package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KbcMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.267 33.758v.043a3.464 3.464 0 0 1-3.464 3.464h0a3.464 3.464 0 0 1-3.464-3.464v-3.53a3.464 3.464 0 0 1 3.464-3.464h0a3.464 3.464 0 0 1 3.464 3.465v.042m-23.312 6.951l-4.306-5.229l4.306-5.193m-5.621-.036v10.458m14.986-5.229h-4.314m4.314 0a2.614 2.614 0 1 1 0 5.23h-4.314v-10.46h4.314a2.614 2.614 0 1 1 0 5.23Zm4.267-18.575a6.14 6.14 0 0 1-6.14 6.14h0a6.14 6.14 0 0 1-6.14-6.14h0a6.14 6.14 0 0 1 6.14-6.14h0a6.14 6.14 0 0 1 6.14 6.14ZM17.24 19.98c4.76 3.475 9.329 1.674 11.956-.88m-11.964.88c-5.278.52-9.88 1.18-12.748 1.944v2.885h39v-6.072h0c-4.208-.03-9.22.1-14.294.363"/>`),
		g.Group(children),
	)
}