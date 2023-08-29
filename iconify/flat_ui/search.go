package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Search(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#2C3E50" fill-rule="evenodd" d="M39.241 64.758a6.025 6.025 0 0 0-8.519 0L8.764 86.717a6.025 6.025 0 0 0 8.519 8.519l21.958-21.958a6.024 6.024 0 0 0 0-8.52z" clip-rule="evenodd"/><path fill="#BDC2C7" fill-rule="evenodd" d="M50.11 53.883a3.015 3.015 0 0 0-4.263 0L34.883 64.848a3.015 3.015 0 0 0 4.263 4.263L50.11 58.146a3.013 3.013 0 0 0 0-4.263z" clip-rule="evenodd"/><path fill="#BDC2C7" d="M100 35.895a8.895 8.895 0 0 1-8.895 8.895H8.895a8.896 8.896 0 0 1 0-17.79h82.209A8.896 8.896 0 0 1 100 35.895z"/><circle cx="9" cy="36" r="4" fill="#D4D7DA"/><path fill="#ECF0F1" d="M68 65c-16.542 0-30-13.458-30-30S51.458 5 68 5s30 13.458 30 30s-13.458 30-30 30z"/><path fill="#2C3E50" d="M68 7c15.439 0 28 12.561 28 28S83.439 63 68 63S40 50.439 40 35S52.561 7 68 7m0-4C50.327 3 36 17.327 36 35s14.327 32 32 32s32-14.327 32-32S85.673 3 68 3z"/><path fill="none" stroke="#fff" stroke-dasharray="0 9 0" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="4" d="M68 13c12.15 0 22 9.85 22 22" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}