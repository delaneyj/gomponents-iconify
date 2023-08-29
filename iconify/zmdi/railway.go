package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Railway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 315V91q0-27 12.5-45t38-26.5t53-11.5t67-3t67 3t53 11.5t38 26.5T341 91v224q0 31-21.5 52.5T267 389l32 32v11H43v-11l32-32q-31 0-53-21.5T0 315zm170.5 32q17.5 0 30-12.5T213 304t-12.5-30.5t-30-12.5t-30 12.5T128 304t12.5 30.5t30 12.5zM299 197V91H43v106h256z"/>`),
		g.Group(children),
	)
}