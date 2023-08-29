package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyesNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M48 0H0v48h48V0zM24.053 12.048c8.947 0 19.947 12 19.947 12s-11 12-19.947 12c-8.948 0-20.053-12-20.053-12s11.105-12 20.053-12zm-17.242 12a59.368 59.368 0 0 1 5.08-4.415a43.176 43.176 0 0 1 3.518-2.456A10.954 10.954 0 0 0 13 24.047c0 2.6.902 4.989 2.41 6.872a43.151 43.151 0 0 1-3.518-2.456a59.368 59.368 0 0 1-5.081-4.415zm28.189 0c0-2.643-.932-5.068-2.485-6.965a42.227 42.227 0 0 1 3.64 2.545a58.582 58.582 0 0 1 5.047 4.42l-.446.433a58.582 58.582 0 0 1-4.6 3.986a42.227 42.227 0 0 1-3.641 2.546A10.955 10.955 0 0 0 35 24.048zm-11 9a9 9 0 1 0 0-18a9 9 0 0 0 0 18zm5-9a5 5 0 1 1-1.748-3.798a1.5 1.5 0 1 0 .547.547A4.98 4.98 0 0 1 29 24.046z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}