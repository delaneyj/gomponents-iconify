package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 8a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2Zm-2 20H4V8h28v20Z" class="clr-i-outline clr-i-outline-path-1"/><circle cx="6.21" cy="10.25" r="1.25" fill="currentColor" class="clr-i-outline clr-i-outline-path-2"/><circle cx="29.81" cy="10.25" r="1.25" fill="currentColor" class="clr-i-outline clr-i-outline-path-3"/><circle cx="6.21" cy="25.42" r="1.25" fill="currentColor" class="clr-i-outline clr-i-outline-path-4"/><circle cx="29.81" cy="25.42" r="1.25" fill="currentColor" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M11.88 18.08a3.59 3.59 0 1 0 3.59-3.59a3.84 3.84 0 0 0-.91.13l.44 1.54a2.08 2.08 0 0 1 .5-.07a2 2 0 1 1-2 2a1.64 1.64 0 0 1 .08-.5L12 17.16a3.53 3.53 0 0 0-.12.92Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="currentColor" d="M15.47 25.73a7.66 7.66 0 0 1-7.65-7.65a7.55 7.55 0 0 1 .27-2l-1.55-.38a9.24 9.24 0 0 0 17.8 4.95h-1.68a7.64 7.64 0 0 1-7.19 5.08Z" class="clr-i-outline clr-i-outline-path-7"/><path fill="currentColor" d="M28.22 17.83a.8.8 0 0 0-.8-.8h-2.76a9.26 9.26 0 0 0-9.19-8.2a9.36 9.36 0 0 0-2.38.32l.42 1.54a7.86 7.86 0 0 1 2-.26A7.66 7.66 0 0 1 23 17h-2.08a.8.8 0 0 0 0 1.6h6.5a.8.8 0 0 0 .8-.77Z" class="clr-i-outline clr-i-outline-path-8"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}