package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M277 224q0-37-17-83.5t-50-83T139 21T67.5 57.5t-50.5 83T0 224q0 58 40.5 98.5T139 363q57 0 97.5-40.5T277 224zm-234 0q0-47 31.5-103.5T139 64t64.5 56.5T235 224q0 40-28 68t-68 28t-68-28t-28-68z"/>`),
		g.Group(children),
	)
}