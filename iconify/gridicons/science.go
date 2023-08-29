package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Science(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.211 19.7c.375.867 1.01 1.3 1.909 1.3h13.76c.898 0 1.535-.433 1.909-1.3c.374-.867.251-1.656-.37-2.367L15.465 11V5h1.668a.917.917 0 0 0 .69-.283a.992.992 0 0 0 .273-.717a.992.992 0 0 0-.273-.717a.917.917 0 0 0-.69-.283H6.867a.916.916 0 0 0-.69.283a.992.992 0 0 0-.272.717c0 .289.091.528.273.717a.916.916 0 0 0 .69.283h1.668v6L3.58 17.333c-.62.711-.743 1.5-.369 2.367zm4.457-4.423h8.664l-2.792-3.544V5h-3.08v6.733l-2.792 3.544z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}