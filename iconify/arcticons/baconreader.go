package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baconreader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.34 7.927c8.791 4.396 7.469 7.43 12.293 11.087c4.49 3.403 9.725 5.329 12.487 8.363a25.314 25.314 0 0 1 4.123 6.146"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.102 4.97c8.792 4.396 7.47 7.43 12.293 11.087c4.49 3.404 9.725 5.33 12.487 8.364a47.685 47.685 0 0 1 4.618 5.99l-4.862 4.646a26.344 26.344 0 0 0-4.124-5.99c-2.762-3.035-7.997-4.96-12.487-8.364c-4.824-3.657-3.5-6.69-12.292-11.086Zm-5.589 9.809c8.078 4.039 7.175 6.451 11.295 10.187c5.253 4.762 7.47 4.289 11.127 7.847a41.448 41.448 0 0 1 4.59 5.342l-5.102 4.874a19.61 19.61 0 0 0-3.806-5.132c-2.78-2.762-5.563-3.034-9.18-6.302c-4.128-3.728-4.859-8.508-12.937-12.547Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.63 41.729a19.608 19.608 0 0 0-3.807-5.133c-2.78-2.762-5.562-3.034-9.18-6.302c-4.127-3.728-4.858-8.508-12.936-12.547"/>`),
		g.Group(children),
	)
}