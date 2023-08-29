package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m175 397l280-280c3-3 27-27 65-40c53-17 107-3 150 40c44 43 57 96 40 149c-12 38-37 64-40 67L369 634c-52 52-182 119-302 0c-119-119-51-250 0-301L390 9c12-12 31-12 43 0s12 31 0 43L111 376c-5 4-106 109 0 215c103 103 204 11 215 0l301-302s17-17 25-41c10-32 3-60-25-88c-60-60-118-10-129 0L218 439c-10 10-17 27 0 44s33 9 43 0l194-194c12-11 32-11 43 0c12 12 12 32 0 44L304 526c-27 26-82 47-129 0c-48-47-27-103 0-129z"/>`),
		g.Group(children),
	)
}