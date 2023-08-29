package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VmwAppLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M28 22h2v8h-8v-2h-2v4h12V20h-4v2z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M14 30H6v-8h2v-2H4v12h12v-4h-2v2z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M8 14H6V6h8v2h2V4H4v12h4v-2z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M20 4v4h2V6h8v8h-2v2h4V4H20z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M11 11h6v6h-6z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M19 11h6v6h-6z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M11 19h6v6h-6z" class="clr-i-outline clr-i-outline-path-7"/><path fill="currentColor" d="M19 19h6v6h-6z" class="clr-i-outline clr-i-outline-path-8"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}