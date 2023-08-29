package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Geocaching(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 5.5h-37v37h37ZM24 42.5v-37M5.5 24h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.613 35.218c3.677-4.073 5.527-.583 8.808-6.05c-4.374.925-6.813-2.187-10.262.379L31.48 42.5M14.1 14.1a13.987 13.987 0 0 0-.999 18.677m19.677-19.676A13.934 13.934 0 0 0 24 10m-8.778 24.898A13.934 13.934 0 0 0 24 38m14-14a13.933 13.933 0 0 0-3.101-8.777"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.611 20.336c0-5.425-9.83-7.654-9.83-12.364"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.1 14.1c.204-.174.296-.324.596-.592c2.368-2.118 4.915-3.18 4.915-5.536"/><circle cx="14.696" cy="9.234" r="1.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}