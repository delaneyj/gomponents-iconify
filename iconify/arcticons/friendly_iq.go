package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FriendlyIq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 28.459c-3.277 0-4.52-1.42-5.097-2.218m-.324-4.199a2.826 2.826 0 0 0-5.651 0m2.825 14.803c1.702.134 1.442 4.592-.732 4.348s-1.797-4.548.732-4.348Z"/><ellipse cx="8.743" cy="7.961" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.44" ry="1.331" transform="rotate(-45 8.743 7.96)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.498 11.902a58.543 58.543 0 0 0-4.03-5.667M12.74 13.6c-1.53-1.697-5.723-3.914-5.723-3.914M24 28.459c3.277 0 4.52-1.42 5.097-2.218m.324-4.199a2.826 2.826 0 0 1 5.651 0M24 43.5c5.097 0 5.673-1.642 5.673-2.707S28.564 36.165 24 36.165s-5.673 3.565-5.673 4.63S18.903 43.5 24 43.5Zm8.247-6.655c-1.702.134-1.442 4.592.732 4.348s1.797-4.548-.732-4.348Z"/><ellipse cx="39.257" cy="7.961" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.331" ry="2.44" transform="rotate(-45 39.257 7.96)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.502 11.902a58.543 58.543 0 0 1 4.03-5.667M35.26 13.6c1.53-1.697 5.723-3.914 5.723-3.914"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.332 8.947h.025A3.447 3.447 0 0 0 29.73 7.34a9.948 9.948 0 0 0-4.17.902A37.064 37.064 0 0 0 28.664 4.5s-6.211 1.553-8.311 4.233h.016C11.754 11.09 6.776 21.472 6.776 27.527c0 6.863 12.56 8.637 17.224 8.637s17.224-1.774 17.224-8.637c0-5.889-4.709-15.867-12.892-18.58Z"/>`),
		g.Group(children),
	)
}