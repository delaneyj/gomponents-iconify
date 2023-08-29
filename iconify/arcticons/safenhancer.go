package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Safenhancer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.678 13.044H24.77c-1.963-.108-5.931-4.238-8.188-4.238H6.68v-.001a2.176 2.176 0 0 0-2.18 2.171v7.307h39v-3.418a1.822 1.822 0 0 0-1.822-1.821Zm1.822 5.239h-39v18.733a2.176 2.176 0 0 0 2.174 2.18h34.645a2.176 2.176 0 0 0 2.181-2.172V18.283Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.546 28.8H27.97v-4.031h-4.03v-1.576a1.454 1.454 0 0 0-2.909 0v1.576H17V28.8h1.576a1.454 1.454 0 1 1 0 2.908H17v4.031h4.031v-1.576a1.454 1.454 0 0 1 2.908 0v1.576h4.031v-4.031h1.576a1.454 1.454 0 1 0 0-2.908Z"/>`),
		g.Group(children),
	)
}