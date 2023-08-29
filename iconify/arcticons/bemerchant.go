package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bemerchant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.562 23.821a7.8 7.8 0 1 1 13.275-8.195a7.8 7.8 0 0 1-13.275 8.196"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.7 26.479a7.8 7.8 0 1 1 3.9-6.755m-15.599 0H39.6m-31.2-.001V5.5m10.232 29.062c0-.887.72-1.607 1.607-1.607m-1.607 0v4.258m-1.957-.81a1.606 1.606 0 0 1-3.003-.796v-1.044a1.607 1.607 0 0 1 3.214 0v.522h-3.214m-8.17 2.121v-6.42l3.213 6.427l3.214-6.417v6.418m12.859-.81a1.606 1.606 0 0 1-3.002-.798v-1.044a1.607 1.607 0 0 1 3-.8m1.548-2.976v6.427m0-2.651a1.607 1.607 0 0 1 3.213 0v2.651m4.76-1.606c0 .887-.719 1.606-1.606 1.606h0a1.606 1.606 0 0 1-1.607-1.606v-1.045c0-.887.72-1.607 1.607-1.607h0c.887 0 1.606.72 1.606 1.607m0 2.651v-4.258m4.761 4.258v-2.65a1.607 1.607 0 0 0-3.214 0m0 2.65v-4.258m5.604-1.325v4.78c0 .444.36.803.803.803h.241m-1.888-4.258h1.688"/>`),
		g.Group(children),
	)
}