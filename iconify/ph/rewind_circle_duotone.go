package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 32a96 96 0 1 0 96 96a96 96 0 0 0-96-96Zm-16 128l-48-32l48-32Zm64 0l-48-32l48-32Z" opacity=".2"/><path d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm51.77-127a8 8 0 0 0-8.21.39l-48 32a8 8 0 0 0 0 13.32l48 32A8 8 0 0 0 176 168a8 8 0 0 0 8-8V96a8 8 0 0 0-4.23-7ZM168 145.05L142.42 128L168 111ZM115.77 89a8 8 0 0 0-8.21.39l-48 32a8 8 0 0 0 0 13.32l48 32A8 8 0 0 0 112 168a8 8 0 0 0 8-8V96a8 8 0 0 0-4.23-7ZM104 145.05L78.42 128L104 111Z"/></g>`),
		g.Group(children),
	)
}