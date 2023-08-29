package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowncoatFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M67 17H5.32v37.804H67V17Z"/><path fill="#FCEA2B" d="M5 30h62v4H5zm0 8h62v4H5z"/><g stroke="#000" stroke-linejoin="round" stroke-width="2"><path fill="none" stroke-linecap="round" d="M67 17H5v38h62V17Z"/><path d="m36 25l2.47 7.946h7.992l-6.466 4.911l2.47 7.947L36 40.893l-6.466 4.91l2.47-7.946l-6.466-4.91h7.992L36 25Z"/></g>`),
		g.Group(children),
	)
}