package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserToUserTransmission(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M14 18C17.866 18 21 14.866 21 11C21 7.13401 17.866 4 14 4C10.134 4 7 7.13401 7 11C7 14.866 10.134 18 14 18Z"/><path fill="#2F88FF" d="M34 18C37.866 18 41 14.866 41 11C41 7.13401 37.866 4 34 4C30.134 4 27 7.13401 27 11C27 14.866 30.134 18 34 18Z"/><path stroke-linecap="round" d="M4 44C4 43.1111 4 40.1111 4 35C4 29.4772 7.77023 25 12.4211 25C14.6667 25 16.3509 25 17.4737 25C21.5587 25 24.0001 29.0269 24.0001 29.0269"/><path stroke-linecap="round" d="M44 44C44 43.1111 44 40.1111 44 35C44 29.4772 40.1849 25 35.4788 25C33.2064 25 31.5022 25 30.366 25C26.4048 25 23.9922 29.0269 24 29.0269"/><path stroke-linecap="round" d="M11 40H38"/><path stroke-linecap="round" d="M34.2954 36.2583L35.535 37.5055L38.0141 39.9998L35.535 42.5613L34.2954 43.8421"/><path stroke-linecap="round" d="M14.3293 36.2321L13.0701 37.4857L10.5517 39.9928L13.0701 42.5415L14.3293 43.8159"/></g>`),
		g.Group(children),
	)
}