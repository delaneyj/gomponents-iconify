package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.004 6.041h-.972L15.017 3l-1.47-3l-1.579 3l.015 3.041H9.962L9.942 3L8.525 0L7.007 3l.023 3.041H4.97L4.941 3L3.524 0L2.006 3l.023 3.041h-.967l.003.979h.967l.012 5.031L1.039 12v1.01h.993V16h2.935v-2.99H7.03V16h2.938v-2.99h2.015V16h3.052l-.02-2.99h.973V12h-.972V7.02h.973l.015-.979zm-11.02.931h2.058v5.079H4.96l.024-5.079zM12.068 12H9.96l-.011-5.072h2.119V12z"/>`),
		g.Group(children),
	)
}