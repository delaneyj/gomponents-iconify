package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pulsar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.961 23.998l-13.463 7.175v-14.35l13.463 7.175Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.977 33.829A19.657 19.657 0 0 1 33.829 6.976m7.194 7.195a19.657 19.657 0 0 1-26.852 26.853M33.829 6.977l8.6-4.477m-36.858 43l8.6-4.476"/>`),
		g.Group(children),
	)
}