package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filepdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 13h.86v-.9h.39c.62 0 1.14-.45 1.14-1.06s-.5-1.05-1.14-1.05H3v3Zm.86-1.59v-.72h.3c.2 0 .37.13.37.35s-.16.36-.37.36h-.3ZM6.19 13h1.19c1 0 1.62-.59 1.62-1.52C9 10.61 8.38 10 7.38 10H6.19v3Zm.86-.71V10.7h.29c.33 0 .78.16.78.78c0 .65-.45.81-.78.81h-.29ZM10 13h.86v-1.07h1.06v-.69h-1.06v-.54h1.21v-.69h-2.06v3Z"/><path fill="currentColor" d="M12.5 16h-10c-.83 0-1.5-.67-1.5-1.5v-13C1 .67 1.67 0 2.5 0h7.09c.4 0 .78.16 1.06.44l2.91 2.91c.28.28.44.66.44 1.06V14.5c0 .83-.67 1.5-1.5 1.5ZM2.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5V4.41a.47.47 0 0 0-.15-.35L9.94 1.15A.51.51 0 0 0 9.59 1H2.5Z"/><path fill="currentColor" d="M13.38 5h-2.91C9.66 5 9 4.34 9 3.53V.62c0-.28.22-.5.5-.5s.5.22.5.5v2.91c0 .26.21.47.47.47h2.91c.28 0 .5.22.5.5s-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}