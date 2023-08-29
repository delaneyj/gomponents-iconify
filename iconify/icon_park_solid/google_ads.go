package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGoogleAds0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" d="M41.355 34.153L29.522 8.776a6 6 0 0 0-10.876 5.072L30.48 39.224a6 6 0 1 0 10.876-5.071Z"/><path stroke-linecap="round" d="M23.438 26.536L17.52 39.224a6 6 0 0 1-7.974 2.902v0a6 6 0 0 1-2.902-7.973L18.374 9"/><circle cx="12.083" cy="36.688" r="6" fill="#fff" transform="rotate(25 12.083 36.688)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGoogleAds0)"/>`),
		g.Group(children),
	)
}