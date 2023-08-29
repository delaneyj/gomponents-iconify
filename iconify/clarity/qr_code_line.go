package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M5.6 4A1.6 1.6 0 0 0 4 5.6V12h8V4Zm4.4 6H6V6h4Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M4 30.4A1.6 1.6 0 0 0 5.6 32H12v-8H4ZM6 26h4v4H6Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M24 32h6.4a1.6 1.6 0 0 0 1.6-1.6V24h-8Zm2-6h4v4h-4Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M30.4 4H24v8h8V5.6A1.6 1.6 0 0 0 30.4 4Zm-.4 6h-4V6h4Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M20 10V8h-4v4h2v-2h2z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M12 12h2v2h-2z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M14 14h4v2h-4z" class="clr-i-outline clr-i-outline-path-7"/><path fill="currentColor" d="M20 6v2h2V4h-8v4h2V6h4z" class="clr-i-outline clr-i-outline-path-8"/><path fill="currentColor" d="M4 14h2v4H4z" class="clr-i-outline clr-i-outline-path-9"/><path fill="currentColor" d="M12 16v2h-2v-4H8v4H6v2H4v2h4v-2h2v2h2v-2h2v-4h-2z" class="clr-i-outline clr-i-outline-path-10"/><path fill="currentColor" d="M20 16h2v2h2v-2h2v-2h-4v-4h-2v2h-2v2h2v2z" class="clr-i-outline clr-i-outline-path-11"/><path fill="currentColor" d="M18 30h-4v2h8v-2h-2v-2h-2v2z" class="clr-i-outline clr-i-outline-path-12"/><path fill="currentColor" d="M22 20v-2h-2v-2h-2v2h-2v2h2v2h2v-2h2z" class="clr-i-outline clr-i-outline-path-13"/><path fill="currentColor" d="M30 20h2v2h-2z" class="clr-i-outline clr-i-outline-path-14"/><path fill="currentColor" d="M22 20h6v2h-6z" class="clr-i-outline clr-i-outline-path-15"/><path fill="currentColor" d="M30 14h-2v2h-2v2h2v2h2v-2h2v-2h-2v-2z" class="clr-i-outline clr-i-outline-path-16"/><path fill="currentColor" d="M20 22h2v6h-2z" class="clr-i-outline clr-i-outline-path-17"/><path fill="currentColor" d="M14 28h2v-2h2v-2h-2v-4h-2v8z" class="clr-i-outline clr-i-outline-path-18"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}