package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 352c-35.35 0-64 28.65-64 64s28.65 64 64 64s64-28.65 64-64s-28.65-64-64-64zm0-192c35.35 0 64-28.65 64-64s-28.65-64-64-64s-64 28.65-64 64s28.65 64 64 64zm192 48H32c-17.67 0-32 14.33-32 32v32c0 17.67 14.33 32 32 32h384c17.67 0 32-14.33 32-32v-32c0-17.67-14.33-32-32-32z"/>`),
		g.Group(children),
	)
}