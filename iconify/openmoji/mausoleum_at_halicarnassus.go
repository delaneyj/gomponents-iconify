package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MausoleumAtHalicarnassus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9b9b9a" d="M27.504 27.245h16.992v13.25H27.504z"/><path fill="#d0cfce" d="M20.687 24.142h30.627v3.103H20.687z"/><path fill="#d0cfce" d="m49.443 24.142l-9.392-9.391h-8.129l-9.365 9.364l26.886.027zm4.299 18.407v13.807H18.258V42.549l2.054-2.054h31.376l2.054 2.054zm-.003 13.812l1.501 1.501v2.995H16.76v-2.995l1.501-1.501h35.478z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.687 24.142h30.627v3.103H20.687z"/><path d="m49.443 24.142l-9.392-9.391h-8.129l-9.365 9.364l26.886.027zM21.84 27.245v13.25m5.664-13.25v13.25m5.664-13.25v13.25m5.664-13.25v13.25m5.664-13.25v13.25m5.664-13.25v13.25m3.582 2.054v13.807H18.258V42.549l2.054-2.054h31.376l2.054 2.054zm-.003 13.812l1.501 1.501v2.995H16.76v-2.995l1.501-1.501h35.478zM36 11.373V7.819m-4.078 6.932l-.962-4.782m9.091 4.782l.962-4.782m-7.002 2.779V9.999m3.978 0v2.749"/></g>`),
		g.Group(children),
	)
}