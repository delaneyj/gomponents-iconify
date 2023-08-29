package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Texture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M352 2q24 6 31 30L32 382q-11-3-19-11T2 352zM189 0h61L0 250v-61zM43 0h42L0 85V43q0-18 12.5-30.5T43 0zm298 384h-42l85-85v42q0 18-13 30q-12 13-30 13zm-207 0l250-250v61L195 384h-61z"/>`),
		g.Group(children),
	)
}