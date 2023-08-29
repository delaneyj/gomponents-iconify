package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.788 17.901v9.077"/><path fill="currentColor" d="M43.538 29.386a.75.75 0 0 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.949 18.005V28.7a.925.925 0 0 1-.928.927H5.427A.925.925 0 0 1 4.5 28.7V18.005m6.428 11.621V17.972l3.221 11.654v-11.62m3.054-.001v11.62h2.838m3.481-11.709h1.847c.48 0 .865.32.865.717v10.191c0 .397-.386.717-.865.717h-1.847c-.48 0-.866-.32-.866-.717v-10.19c0-.398.387-.718.866-.718Zm9.388 1.392v-.675c0-.397-.386-.717-.865-.717h-1.847c-.48 0-.866.32-.866.717v10.191c0 .397.386.717.866.717h1.846c.48 0 .866-.32.866-.717v-.797m2.885-10.111v5.813m2.101 0l1.678-5.827m0 11.654l-1.678-5.827h-2.1v5.813"/>`),
		g.Group(children),
	)
}