package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 5H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2ZM4 29V7h28v22Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M20.971 11.243c2.343 2.343 1.393 7.092-2.122 10.606c-3.515 3.515-8.263 4.465-10.606 2.121c-2.344-2.343-1.394-7.092 2.121-10.606c3.515-3.515 8.264-4.464 10.607-2.121Zm-9.335 3.394c-2.812 2.812-3.761 6.421-2.121 8.061c1.64 1.64 5.249.691 8.061-2.121c2.812-2.812 3.762-6.421 2.121-8.061c-1.64-1.64-5.249-.691-8.061 2.121Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M28 22a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-3-1.4a1.4 1.4 0 1 0 0 2.8a1.4 1.4 0 0 0 0-2.8Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}