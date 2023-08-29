package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.47 3.84a5.78 5.78 0 0 0-7.37-.63a16.08 16.08 0 0 1 8.2 7.65a5.73 5.73 0 0 0-.83-7.02Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M25.33 21.54a.9.9 0 0 0-.41-1.2l-3.2-1.56L24.89 22a.89.89 0 0 0 .44-.46Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M18 8.6a.9.9 0 0 0-.9.9v4.6l1.8 1.81V9.5a.9.9 0 0 0-.9-.9Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M11.42 3.43a5.8 5.8 0 0 0-5.81-.81l2.69 2.7a16 16 0 0 1 3.12-1.89Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M18 4a13.91 13.91 0 0 0-8.3 2.75l1.42 1.43a12 12 0 0 1 16.7 16.72l1.42 1.43A14 14 0 0 0 18 4Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="m1.56 4.21l1.17 1.17a5.7 5.7 0 0 0 .67 6.1a15.78 15.78 0 0 1 2.06-3.36l1.42 1.43a13.94 13.94 0 0 0 1.23 18.33l-2.55 2.55A1 1 0 1 0 7 31.84l2.66-2.66a13.89 13.89 0 0 0 16.8 0l4.14 4.15L32 31.9L3 2.8ZM25 27.72A11.89 11.89 0 0 1 18 30A12 12 0 0 1 6 18a11.89 11.89 0 0 1 2.29-7Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}