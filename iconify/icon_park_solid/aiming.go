package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aiming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAiming0"><g fill="none"><circle cx="24" cy="24" r="20" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path fill="#fff" fill-rule="evenodd" d="M24 37v7v-7Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 37v7"/><path fill="#fff" fill-rule="evenodd" d="M36 24h8h-8Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 24h8"/><path fill="#fff" fill-rule="evenodd" d="M4 24h7h-7Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 24h7"/><path fill="#fff" fill-rule="evenodd" d="M24 11V4v7Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 11V4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAiming0)"/>`),
		g.Group(children),
	)
}