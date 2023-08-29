package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="svgIDa"><g fill="none"><path fill="#fff" fill-rule="evenodd" d="M24 3v3.15V3Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 3v3.15"/><path fill="#fff" fill-rule="evenodd" d="m38.85 9.15l-2.228 2.228l2.227-2.227Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m38.85 9.15l-2.228 2.228"/><path fill="#fff" fill-rule="evenodd" d="M45 24h-3.15H45Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M45 24h-3.15"/><path fill="#fff" fill-rule="evenodd" d="m38.85 38.85l-2.228-2.228l2.227 2.227Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m38.85 38.85l-2.228-2.228"/><path fill="#fff" fill-rule="evenodd" d="M24 45v-3.15V45Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 45v-3.15"/><path fill="#fff" fill-rule="evenodd" d="m9.15 38.85l2.228-2.228l-2.227 2.227Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m9.15 38.85l2.228-2.228"/><path fill="#fff" fill-rule="evenodd" d="M3 24h3.15H3Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M3 24h3.15"/><path fill="#fff" fill-rule="evenodd" d="m9.15 9.15l2.228 2.228l-2.227-2.227Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m9.15 9.15l2.228 2.228"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 36c6.627 0 12-5.373 12-12s-5.373-12-12-12s-12 5.373-12 12s5.373 12 12 12Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#svgIDa)"/>`),
		g.Group(children),
	)
}