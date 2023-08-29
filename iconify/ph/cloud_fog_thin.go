package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudFogThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 196H72a4 4 0 0 1 0-8h48a4 4 0 0 1 0 8Zm64-8h-24a4 4 0 0 0 0 8h24a4 4 0 0 0 0-8Zm-24 32h-56a4 4 0 0 0 0 8h56a4 4 0 0 0 0-8Zm68-128a72.08 72.08 0 0 1-72 72H76a48 48 0 1 1 11.51-94.61A72.08 72.08 0 0 1 228 92Zm-8 0a64.06 64.06 0 0 0-128-3.77a4 4 0 0 1-8-.46a71.63 71.63 0 0 1 1.42-10.65A40 40 0 1 0 76 156h80a64.07 64.07 0 0 0 64-64Z"/>`),
		g.Group(children),
	)
}