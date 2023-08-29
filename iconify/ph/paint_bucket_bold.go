package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintBucketBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M231.5 162.63a12 12 0 0 0-15 0C214 164.64 192 183 192 208a32 32 0 0 0 64 0c0-25-22-43.36-24.5-45.37ZM224 216a8 8 0 0 1-8-8c0-6.8 3.91-13.64 8-18.83c4.07 5.17 8 12 8 18.83a8 8 0 0 1-8 8Zm8.49-101L125 7.51a12 12 0 0 0-17 0L70.25 45.29L48.48 23.51a12 12 0 1 0-17 17l21.8 21.75l-41.08 41.08a28 28 0 0 0 0 39.6l84.86 84.86a28 28 0 0 0 39.6 0l95.83-95.8a12 12 0 0 0 0-17Zm-112.8 95.83a4 4 0 0 1-5.66 0L29.17 126a4 4 0 0 1 0-5.66l41.08-41.11l21.27 21.26a36 36 0 1 0 17-17L87.23 62.26L116.52 33L207 123.48ZM124 104a12 12 0 1 1-12 12a12 12 0 0 1 12-12Z"/>`),
		g.Group(children),
	)
}