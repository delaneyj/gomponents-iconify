package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFalklandIslands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#fff" d="M9.887 18H6v2.332L32.113 36H36v-2.332L9.887 18z"/><path fill="#fff" d="M36 20.332V18h-3.887L6 33.668V36h3.887L36 20.332z"/><path fill="#fff" d="M6 24h30v6H6z"/><path fill="#fff" d="M18 18h6v18h-6z"/><path fill="#d22f27" d="M20 18h2v18h-2z"/><path fill="#d22f27" d="M6 26h30v2H6zm30 7.668L29.887 30H26l10 6v-2.332zM36 18h-3.887L24 22.868V24h2l10-6zM6 20.332L12.113 24H16L6 18v2.332zM6 36h3.887L18 31.132V30h-2L6 36z"/><path fill="#61b2e4" d="M59 29v5c0 9-2.667 13-8 15c-5.333-2-8-6-8-15v-5Z"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M59 29v5c0 9-2.667 13-8 15c-5.333-2-8-6-8-15v-5Z"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M49 33h5v2h-5zm5 2v2m-5-2v2m-2-4h2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M44.035 40.115c1 .433 1.91.81 2.965.885c2 0 2-1 4-1"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M45.721 44.814c.421.164.858.157 1.279.186c2 0 2-1 4-1"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M57.965 40.115c-1 .433-1.91.81-2.965.885c-2 0-2-1-4-1"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M56.279 44.814c-.421.164-.858.157-1.279.186c-2 0-2-1-4-1"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}