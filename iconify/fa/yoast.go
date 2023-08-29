package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yoast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M339 218h691l-26 72H339q-110 0-188.5 79T72 558v771q0 95 60.5 169.5T286 1592q23 5 98 5v72h-45q-140 0-239.5-100T0 1329V558q0-140 99.5-240T339 218zM1190 0h247L955 1294q-23 61-40.5 103.5t-45 98t-54 93.5t-64.5 78.5t-79.5 65t-95.5 41t-116 18.5v-195q163-26 220-182q20-52 20-105q0-54-20-106L395 471h228l187 585zm474 558v1111H869q37-55 45-73h678V558q0-85-49.5-155T1413 304l25-67q101 34 163.5 123.5T1664 558z"/>`),
		g.Group(children),
	)
}