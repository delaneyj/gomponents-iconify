package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnsChangerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.738 13.614C5.76 12.17 11.74 4.855 24 4.854c12.26.001 18.24 7.315 19.262 8.76c1.97 16.92-8.615 26.438-19.262 29.532C13.353 40.053 2.768 30.535 4.738 13.614Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.762 11.107l2.923-.462l1.277-2.67l1.342 2.637l2.934.39l-2.093 2.091l.537 2.911l-2.637-1.345l-2.602 1.41l.464-2.923l-2.145-2.039ZM13.476 21.24v8.282h1.864c1.967 0 3.623-1.657 3.623-3.624v-1.035c0-1.967-1.656-3.623-3.623-3.623h-1.864Zm7.78 8.282v-8.283l5.488 8.283v-8.283m2.396 7.351c.518.621 1.139.932 2.07.932h1.243c1.139 0 2.07-.932 2.07-2.07s-.931-2.071-2.07-2.071h-1.346c-1.139 0-2.07-.932-2.07-2.07s.931-2.071 2.07-2.071h1.243c.931 0 1.552.207 2.07.931"/>`),
		g.Group(children),
	)
}