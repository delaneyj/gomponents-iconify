package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LarkPlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="5.044" height="13.102" x="3.5" y="20.485" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.159"/><circle cx="14.849" cy="28.362" r="1.855" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="5.044" height="13.102" x="39.456" y="20.485" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.159" transform="rotate(-180 41.978 27.036)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.932 6.692a18.381 18.381 0 0 1 16.86 13.756M24 37.034c2.518 0 4.213-3.448 4.213-3.448S26.106 30.45 24 30.45s-4.213 3.136-4.213 3.136s1.695 3.448 4.213 3.448Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.213 33.586A7.682 7.682 0 0 1 24 35.24a7.682 7.682 0 0 1-4.213-1.654"/><circle cx="33.151" cy="28.362" r="1.855" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.234 20.988c-2.288-6.514-8.342-10.441-12.706-11.974l.005-.001A13.842 13.842 0 0 0 13.85 4.73s5.077.655 7.632 3.275a12.038 12.038 0 0 0-10.973 1.008s4.762-1.524 8.605.535c-4.182 1.813-9.289 5.579-11.348 11.439m.297 11.791C11.329 40.544 20.264 43.642 24 43.642s12.67-3.097 15.936-10.863"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.875 8.559a18.412 18.412 0 0 0-9.668 11.89"/>`),
		g.Group(children),
	)
}