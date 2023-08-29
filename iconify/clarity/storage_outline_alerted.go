package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M19.51 9.09L18 9.11A37.6 37.6 0 0 1 7 7.76v2.09a43.53 43.53 0 0 0 11 1.27h.61A3.66 3.66 0 0 1 19 9.89Z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M28.83 15.4A38.37 38.37 0 0 1 18 16.7a37.45 37.45 0 0 1-11-1.36v2.08a43.33 43.33 0 0 0 11 1.28c4 0 9.93-.48 13-2v5.17c-.33.86-5.06 2.45-13 2.45a37.45 37.45 0 0 1-11-1.4V25a43.33 43.33 0 0 0 11 1.28c4 0 9.93-.48 13-2v5.1c-.35.86-5.08 2.45-13 2.45S5.3 30.2 5 29.37V6.82c.3-.82 5-2.46 13-2.46c1.5 0 2.89.06 4.15.16l1.1-1.9c-1.86-.18-3.7-.26-5.25-.26c-5.57 0-15 .93-15 4.43v22.58c0 3.49 9.43 4.43 15 4.43s15-.93 15-4.43v-14Z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-3--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}