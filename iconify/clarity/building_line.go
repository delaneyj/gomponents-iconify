package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31 8h-8v2h8v21h-8v2h10V10a2 2 0 0 0-2-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M19.88 3H6.12A2.12 2.12 0 0 0 4 5.12V33h18V5.12A2.12 2.12 0 0 0 19.88 3ZM20 31h-3v-3H9v3H6V5.12A.12.12 0 0 1 6.12 5h13.76a.12.12 0 0 1 .12.12Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M8 8h2v2H8z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M12 8h2v2h-2z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M16 8h2v2h-2z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M8 13h2v2H8z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M12 13h2v2h-2z" class="clr-i-outline clr-i-outline-path-7"/><path fill="currentColor" d="M16 13h2v2h-2z" class="clr-i-outline clr-i-outline-path-8"/><path fill="currentColor" d="M8 18h2v2H8z" class="clr-i-outline clr-i-outline-path-9"/><path fill="currentColor" d="M12 18h2v2h-2z" class="clr-i-outline clr-i-outline-path-10"/><path fill="currentColor" d="M16 18h2v2h-2z" class="clr-i-outline clr-i-outline-path-11"/><path fill="currentColor" d="M8 23h2v2H8z" class="clr-i-outline clr-i-outline-path-12"/><path fill="currentColor" d="M12 23h2v2h-2z" class="clr-i-outline clr-i-outline-path-13"/><path fill="currentColor" d="M16 23h2v2h-2z" class="clr-i-outline clr-i-outline-path-14"/><path fill="currentColor" d="M23 13h2v2h-2z" class="clr-i-outline clr-i-outline-path-15"/><path fill="currentColor" d="M27 13h2v2h-2z" class="clr-i-outline clr-i-outline-path-16"/><path fill="currentColor" d="M23 18h2v2h-2z" class="clr-i-outline clr-i-outline-path-17"/><path fill="currentColor" d="M27 18h2v2h-2z" class="clr-i-outline clr-i-outline-path-18"/><path fill="currentColor" d="M23 23h2v2h-2z" class="clr-i-outline clr-i-outline-path-19"/><path fill="currentColor" d="M27 23h2v2h-2z" class="clr-i-outline clr-i-outline-path-20"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}