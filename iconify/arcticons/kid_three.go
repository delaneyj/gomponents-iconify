package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KidThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.605 18.128a1.882 1.882 0 0 1 1.84.875a4.63 4.63 0 0 1 .828 1.735c.43 1.704-.068 3.3-1.112 3.563s-2.239-.905-2.668-2.61s.068-3.3 1.112-3.563Zm-4.667 9.865l25.391-6.932l2.31 8.461L43 34.511l-25.389 6.931l-3.673-13.449z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.83 9.12l-8.034 9.537a7.615 7.615 0 0 0 4.088 12.327l3.874.889m13.635-7.84a15.486 15.486 0 0 0-.453-4.606c-1.36-5.368-5.123-9.067-8.415-8.224s-4.869 5.858-3.509 11.19a14.97 14.97 0 0 0 1.887 4.533"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.55 23.206a22.29 22.29 0 0 0-.664-5.111c-1.922-7.43-7.135-12.521-11.683-11.376s-6.726 8.105-4.856 15.526a21.63 21.63 0 0 0 1.983 5.11m6.076 7.739a2.84 2.84 0 0 1 .643-.485a1.52 1.52 0 0 1 1.808.04c.554.554.355 1.652-.445 2.451s-1.897 1-2.45.445s-.355-1.651.444-2.45Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.817 31.825c-1.962-3.46-3.167-4.651-5.293-3.905l2.509 6.968m4.912.256l8.637-2.361m-8.817-17.969c.166-.046.359-.101.574-.164"/><path fill="none" stroke="currentColor" stroke-dasharray="1.266 2.11" stroke-linecap="round" stroke-linejoin="round" d="M31.377 14.103c2.03-.469 4.384-.712 5.08.71c1.057 2.158-3.681 6.097-2.573 9.006a2.619 2.619 0 0 0 2.52 1.862c.983-.1 1.804-1.143 1.772-3.654a3.438 3.438 0 0 0-1.351-2.59"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.983 18.795a21.806 21.806 0 0 0-.5-.33"/>`),
		g.Group(children),
	)
}