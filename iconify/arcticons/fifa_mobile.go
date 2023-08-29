package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FifaMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.393 26.252l6.982-13.48l2.117 13.48m-.695-4.55H32.73m-9.036 4.55l2.432-13.48h6.74m-7.956 6.74h4.38m-20.783 6.74l2.432-13.48h6.74m-7.956 6.74h4.38m7.8-6.74l-2.432 13.48m-.448 8.977h0a2.114 2.114 0 0 1-2.13-2.1V30.89c0-1.23.984-2.212 2.13-2.212h0a2.202 2.202 0 0 1 2.211 2.192v2.149a2.202 2.202 0 0 1-2.193 2.21h-.018Zm6.92-3.275H23.16m2.621 0c.9 0 1.638.736 1.638 1.637c0 .9-.737 1.637-1.638 1.637H23.08v-6.55h2.702c.9 0 1.638.736 1.638 1.636a1.642 1.642 0 0 1-1.638 1.639v0ZM39.5 35.228h-3.276v-6.55H39.5m-3.275 3.275h2.13M8.5 35.147v-6.47l3.276 6.551l3.276-6.55v6.55m14.21-6.55v6.55m1.843-6.55v6.55h3.276"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}