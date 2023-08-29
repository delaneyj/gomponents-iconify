package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeRestoreSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 192q0 18-12.5 30.5T256 235t-30.5-12.5T213 192t12.5-30.5T256 149t30.5 12.5T299 192zM256 0q80 0 136 56t56 136t-56 136t-136 56q-65 0-117-40l30-30q40 27 87 27q62 0 105.5-43.5T405 192T361.5 86.5T256 43T150.5 86.5T107 192h64l-86 85l-85-85h64q0-80 56-136T256 0z"/>`),
		g.Group(children),
	)
}