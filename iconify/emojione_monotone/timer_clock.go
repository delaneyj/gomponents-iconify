package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M31.999 3C15.431 3 2 18.711 2 38.094C2 57.475 15.431 61 31.999 61S62 57.475 62 38.094C62 18.711 48.567 3 31.999 3zM32 50.152c-12.416 0-22.479-10.215-22.479-22.816S19.584 4.52 32 4.52c12.414 0 22.479 10.215 22.479 22.816S44.414 50.152 32 50.152z"/><ellipse cx="22.404" cy="43.955" fill="currentColor" rx="1.308" ry="1.289" transform="rotate(-59.987 22.407 43.957)"/><ellipse cx="41.595" cy="10.715" fill="currentColor" rx="1.308" ry="1.289" transform="rotate(119.993 41.595 10.714)"/><ellipse cx="15.379" cy="36.932" fill="currentColor" rx="1.29" ry="1.307" transform="rotate(-119.98 15.38 36.932)"/><ellipse cx="48.62" cy="17.74" fill="currentColor" rx="1.308" ry="1.289" transform="rotate(149.979 48.621 17.741)"/><ellipse cx="12.808" cy="27.336" fill="currentColor" rx="1.308" ry="1.289"/><ellipse cx="51.191" cy="27.336" fill="currentColor" rx="1.309" ry="1.289"/><ellipse cx="15.379" cy="17.739" fill="currentColor" rx="1.289" ry="1.31" transform="rotate(120.006 15.379 17.738)"/><ellipse cx="48.621" cy="36.931" fill="currentColor" rx="1.289" ry="1.308" transform="rotate(-59.979 48.622 36.932)"/><ellipse cx="22.404" cy="10.715" fill="currentColor" rx="1.308" ry="1.289" transform="rotate(59.974 22.403 10.714)"/><path fill="currentColor" d="M40.941 42.822a1.3 1.3 0 0 0-.463 1.779a1.297 1.297 0 0 0 1.771.486c.615-.354.824-1.15.461-1.775a1.298 1.298 0 0 0-1.769-.49"/><ellipse cx="32" cy="8.145" fill="currentColor" rx="1.289" ry="1.309"/><ellipse cx="32" cy="46.527" fill="currentColor" rx="1.289" ry="1.309"/><path fill="currentColor" d="M33.484 11.411c7.32.743 13.033 6.926 13.033 14.442c0 8.018-6.5 14.518-14.519 14.518s-14.518-6.5-14.518-14.518c0-7.517 5.712-13.699 13.032-14.442C22.375 12.161 16 19.001 16 27.336c0 8.836 7.162 16 15.999 16S48 36.172 48 27.336c0-8.335-6.376-15.175-14.516-15.925"/><path fill="currentColor" d="M32 11.336c-2.721 0-4.926 14.337-4.926 21.787c0 7.448 9.85 7.448 9.85 0c0-7.45-2.204-21.787-4.924-21.787"/>`),
		g.Group(children),
	)
}