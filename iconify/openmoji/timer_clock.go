package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiTimerClock0" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M12.605 57.682v-23c0-12.702 10.297-23 23-23s23 10.298 23 23v23h-46z"/></defs><g stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10"><path fill="#9b9b9a" d="M12.605 57.682v-23c0-12.702 10.297-23 23-23s23 10.298 23 23v23h-46z"/><circle cx="35.605" cy="34.682" r="15" fill="#d0cfce"/><circle cx="35.605" cy="34.682" r="15" fill="#FFF"/><path fill="#FFF" d="M35.417 35.087h-6.812"/></g><use href="#openmojiTimerClock0" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><circle cx="35.605" cy="34.682" r="15" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><circle cx="35.605" cy="34.682" r="15" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><use href="#openmojiTimerClock0" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><circle cx="35.605" cy="34.682" r="15" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="1.758" d="M35.417 35.087h-6.812"/><circle cx="24.407" cy="34.988" r="1"/><circle cx="46.678" cy="34.988" r="1"/><circle cx="27.668" cy="42.862" r="1" transform="rotate(-45.001 27.668 42.862)"/><circle cx="43.416" cy="27.114" r="1" transform="rotate(-45.001 43.416 27.114)"/><circle cx="35.542" cy="46.123" r="1"/><circle cx="35.542" cy="23.852" r="1"/><circle cx="43.416" cy="42.862" r="1" transform="rotate(-45.001 43.416 42.863)"/><circle cx="27.668" cy="27.114" r="1" transform="rotate(-45.001 27.668 27.114)"/>`),
		g.Group(children),
	)
}