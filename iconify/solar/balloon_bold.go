package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalloonBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M19.5 9.56c.033 4.143-3.419 7.973-7.56 7.94c-4.143-.033-7.406-3.918-7.44-8.06A7.355 7.355 0 0 1 11.94 2c4.141.034 7.526 3.419 7.56 7.56Zm-6.994-4.31a.75.75 0 0 0-.012 1.5a2.285 2.285 0 0 1 2.256 2.256a.75.75 0 0 0 1.5-.012a3.785 3.785 0 0 0-3.744-3.744Z" clip-rule="evenodd"/><path d="M14.167 18.214c.332 1.063-.356 2.132-1.417 2.348V22a.75.75 0 0 1-1.5 0v-1.438c-1.06-.216-1.75-1.285-1.417-2.348l.007-.021a7.743 7.743 0 0 0 4.32-.002l.007.023Z"/></g>`),
		g.Group(children),
	)
}