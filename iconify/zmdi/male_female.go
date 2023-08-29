package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaleFemale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M32 429V269H0V152q0-18 12.5-30.5T43 109h64q17 0 29.5 12.5T149 152v117h-32v160H32zm267 0h-64V301h-64l54-162q4-14 15.5-22t24.5-8h3q14 0 25 8t15 22l55 162h-64v128zM74.5 88q-17.5 0-30-12.5T32 45.5t12.5-30T74.5 3t30 12.5t12.5 30t-12.5 30t-30 12.5zm192 0q-17.5 0-30-12.5t-12.5-30t12.5-30t30-12.5t30 12.5t12.5 30t-12.5 30t-30 12.5z"/>`),
		g.Group(children),
	)
}