package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M15 8h9v2h-9z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M15 12h9v2h-9z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M15 16h9v2h-9z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M15 20h9v2h-9z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M15 24h9v2h-9z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M11 8h2v2h-2z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M11 12h2v2h-2z" class="clr-i-outline clr-i-outline-path-7"/><path fill="currentColor" d="M11 16h2v2h-2z" class="clr-i-outline clr-i-outline-path-8"/><path fill="currentColor" d="M11 20h2v2h-2z" class="clr-i-outline clr-i-outline-path-9"/><path fill="currentColor" d="M11 24h2v2h-2z" class="clr-i-outline clr-i-outline-path-10"/><path fill="currentColor" d="M28 2H8a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm0 30H8V4h20Z" class="clr-i-outline clr-i-outline-path-11"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}