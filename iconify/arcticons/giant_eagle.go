package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiantEagle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="18.192" cy="17.915" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.326 20.504v6c0 1.1-.9 2-2 2h0c-.6 0-1.1-.2-1.4-.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.326 20.404h0c1.1 0 2 .9 2 2v1.3c0 1.1-.9 2-2 2h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2Zm4.866-.087v5.4m17.181-7.002v6a.94.94 0 0 0 1 1h.3m-2.4-5.4h2.201m-11.657 3.402c0 1.1-.9 2-2 2h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2h0c1.1 0 2 .9 2 2m0 3.3v-5.4m7.119 5.4v-3.3c0-1.1-.9-2-2-2h0c-1.1 0-2 .9-2 2m0 3.3v-5.4m-.521 10.503v6c0 1.1-.9 2-2 2h0c-.6 0-1.1-.2-1.4-.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.415 30.72h0c1.1 0 2 .9 2 2v1.3c0 1.1-.9 2-2 2h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2Zm4.694-2.693v7a.94.94 0 0 0 1 1h.3m-10.591-2c0 1.1-.9 2-2 2h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2h0c1.1 0 2 .9 2 2m0 3.3v-5.4m16.357 4.4c-.3.6-1 1-1.7 1h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2h0c1.1 0 2 .9 2 2v.7h-4m-19.25 1.6c-.3.6-1 1-1.7 1h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2h0c1.1 0 2 .9 2 2v.7h-4M8.49 14.279h9.97m11 0h10.03M18.981 9.807s-.093 5.855 5.078 5.596c1.018-4.835-5.078-5.596-5.078-5.596Zm5.078 5.596c2.151.008 3.394-.898 4.281-2.521c.621-1.515.862-3.053.497-4.628c-.014.003-4.414.377-5.077 4.355"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}