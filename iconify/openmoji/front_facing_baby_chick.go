package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrontFacingBabyChick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<ellipse cx="36" cy="43" fill="#FCEA2B" stroke="#FCEA2B" stroke-miterlimit="10" stroke-width="1.8" rx="15" ry="11"/><path fill="#FCEA2B" d="M42.25 36c2 9 9 9 9 9c2-6-6-11-6-11c7-6 1-10 1-10c0-10-10.375-11-10.375-11C25.5 15 26.25 24 26.25 24c-6 4 .625 9.75.625 9.75S19.25 38 21.25 45c0 0 8 0 8-9"/><path fill="#F1B31C" d="M36.622 26.603s5.925-.29 0 5.233c0 0-5.925-5.233 0-5.233z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M23 49s4.797 5.025 13.48 5.012C45 54 49 49 49 49M33.625 13.5l-2.5-1.875m4 1.25l-1-2.125m3 2.5L37.25 10m-5.904 43.731l-.971 4.519L27.25 60m6-1l-3-1"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m40.154 53.731l.971 4.519L44.25 60m-6-1l3-1"/><circle cx="32.25" cy="23" r="2"/><circle cx="40.25" cy="23" r="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M36.372 26.603s5.913-.29 0 5.233c0 0-5.913-5.233 0-5.233z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M42 36c2 9 9 9 9 9c2-6-6-11-6-11c7-6 1-10 1-10c0-10-10.375-11-10.375-11C25.25 15 26 24 26 24c-6 4 .625 9.75.625 9.75S19 38 21 45c0 0 8 0 8-9"/>`),
		g.Group(children),
	)
}