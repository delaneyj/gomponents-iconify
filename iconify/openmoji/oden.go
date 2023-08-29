package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="35.999" r="10" fill="#f4aa41" transform="rotate(-45.001 36 36)"/><path fill="#e27022" d="M43.07 28.92c-2.593-2.595-6.257-3.455-9.571-2.602a9.939 9.939 0 0 1 4.571 2.602c3.907 3.905 3.907 10.236.001 14.142a9.949 9.949 0 0 1-4.571 2.602c3.313.854 6.977-.008 9.571-2.602c3.906-3.906 3.906-10.237 0-14.142z"/><path fill="#a57939" d="m19.125 36l-4.127-21l21.001 4.128z"/><path fill="#6A462F" d="M36 19.128L14.999 15l.248 1.263l13.129 3.564l-9.775 13.497l.525 2.676z"/><path fill="#d0cfce" d="M39.582 55.418A9.943 9.943 0 0 0 44.07 58l13.897-13.896a9.95 9.95 0 0 0-7.07-7.072L37 50.928a9.948 9.948 0 0 0 2.582 4.49z"/><path fill="#9b9b9a" d="M58 44.103a9.95 9.95 0 0 0-7.07-7.07l-.076.074c.19.202.385.39.565.61c1.289 1.601 2.142 3.53 2.581 5.568L42.582 57.443c.496.223 1.002.418 1.522.557L58 44.103z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="36.001" cy="35.999" r="10" transform="rotate(-45.001 36.002 36)"/><path d="m54 54l13 13M19.17 36.042l-4.127-20.999l21.001 4.127zm20.515 19.343a9.943 9.943 0 0 0 4.489 2.582L58.07 44.07a9.95 9.95 0 0 0-2.58-4.488A9.928 9.928 0 0 0 51 37L37.103 50.897a9.952 9.952 0 0 0 2.582 4.488z"/></g>`),
		g.Group(children),
	)
}