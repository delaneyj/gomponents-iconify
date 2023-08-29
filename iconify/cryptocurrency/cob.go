package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cob(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm0-7h4.536l-2.272-3.957h-4.528L11.464 25H16zm-2.263-14.044h4.527L20.536 7h-9.072l2.272 3.957zm9.999 9.539L26 16.538h-4.536l-2.265 3.957h.001l2.263 3.958l2.273-3.957h-.002h.002zM8.264 11.496L6 15.453h4.536l2.264-3.957l-2.264-3.957l-2.272 3.957zm4.535 9h.001l-2.264-3.958H6l2.264 3.957l2.272 3.958l2.264-3.957zm10.937-9l-2.273-3.957l-2.264 3.957l2.264 3.957H26l-2.264-3.957z"/>`),
		g.Group(children),
	)
}