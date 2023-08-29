package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myradar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="33.243" cy="24.861" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.695 24.191c0 3.59-5.532 16.644-6.5 16.65c-1.166.008-6.5-13.06-6.5-16.65a6.5 6.5 0 1 1 13 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.036 5.5c-1.961 5.317 1.515 9.074 1.151 14.528M15.65 5.5s-5.576 4.723-4.399 7.313c2.158 4.746-5.741 7.227-5.751 7.203M22.725 5.5s-2.75 5.618-1.944 8.364C23.942 24.631 5.5 27.57 5.5 27.57M30.47 5.5c-2.321 6.472.283 9.343.062 12.76m-3.734 7c.135.34-4.155 3.442-6.782 3.202c-4.965-.454-13.4 6.714-14.516 6.151m23.218-2.668c-7.151 3.449-11.235-.812-20.91 10.555"/>`),
		g.Group(children),
	)
}