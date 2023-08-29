package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M4 4h6v6H4z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M4 15h6v6H4z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M4 26h6v6H4z" class="clr-i-solid clr-i-solid-path-3"/><path fill="currentColor" d="M15 4h6v6h-6z" class="clr-i-solid clr-i-solid-path-4"/><path fill="currentColor" d="M15 15h6v6h-6z" class="clr-i-solid clr-i-solid-path-5"/><path fill="currentColor" d="M15 26h6v6h-6z" class="clr-i-solid clr-i-solid-path-6"/><path fill="currentColor" d="M26 4h6v6h-6z" class="clr-i-solid clr-i-solid-path-7"/><path fill="currentColor" d="M26 15h6v6h-6z" class="clr-i-solid clr-i-solid-path-8"/><path fill="currentColor" d="M26 26h6v6h-6z" class="clr-i-solid clr-i-solid-path-9"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}