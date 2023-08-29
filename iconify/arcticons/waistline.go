package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Waistline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.93 4.76a25.75 25.75 0 0 0-5.62.46a14.43 14.43 0 0 0-4.55 1.71c-.49.45-.51 3.29 0 5.11a9.94 9.94 0 0 0 6.76 6.84a9.77 9.77 0 0 0 10.88-4.41A10.44 10.44 0 0 0 33.7 9c0-1.65 0-1.79-.57-2.21A18.9 18.9 0 0 0 28 5a16.62 16.62 0 0 0-3.08-.28Zm-4.52 19.73"/><circle cx="24.04" cy="8.15" r=".75" fill="currentColor"/><circle cx="30.28" cy="9.65" r=".75" fill="currentColor"/><circle cx="17.8" cy="9.65" r=".75" fill="currentColor"/><circle cx="20.92" cy="8.9" r=".75" fill="currentColor"/><circle cx="27.16" cy="8.9" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.61 33.5c-2.55-1.63-2.55-3.16-2.55-4c0-2.16 3.29-4.45 6.32-4.45s2.27 2.49 1.93 3.41s-2.9 6.73-3.24 7.65s-.88 3 .72 3c3 0 6.9-1.52 7.53-8.71c0 0-5.81 8.54-.52 8.54c5.84 0 6.47-9.13 6.71-13.84"/><circle cx="24.04" cy="15.76" r="1.26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.04 14.5v-3.79M18.26 5.5H7.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2H29.83"/>`),
		g.Group(children),
	)
}