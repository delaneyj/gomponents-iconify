package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.065 29.894c-9.06 0-8.97-13.994.054-14.736c1.655-6.32 6.642-9.082 11.435-8.824a10.9 10.9 0 0 1 9.18 5.869c11.253 1.934 10.098 17.691 0 17.691h-20.67Z"/><circle cx="17.457" cy="32.388" r=".75" fill="currentColor"/><circle cx="28.431" cy="38.364" r=".75" fill="currentColor"/><circle cx="30.495" cy="33.473" r=".75" fill="currentColor"/><circle cx="22.925" cy="35.081" r=".75" fill="currentColor"/><circle cx="26.163" cy="32.388" r=".75" fill="currentColor"/><circle cx="11.961" cy="33.581" r=".75" fill="currentColor"/><circle cx="17.898" cy="36.45" r=".75" fill="currentColor"/><circle cx="37.653" cy="32.775" r=".75" fill="currentColor"/><circle cx="34.493" cy="40.182" r=".75" fill="currentColor"/><circle cx="34.063" cy="35.655" r=".75" fill="currentColor"/><circle cx="14.126" cy="40.932" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}