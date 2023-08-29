package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsaSwimming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.349 27.934l2.017.717m-26.354-.717l-2.017.717m6.886-2.168l-3.635 5.914m8.008-6.562l-1.752 10.593m4.676-8.813l.869 2.676h2.814l-2.276 1.654l.869 2.676l-2.276-1.654l-2.277 1.654l.87-2.676l-2.277-1.654h2.814l.87-2.676z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.18 9.521c5.743 3.88 12.056 1.103 14.186-.57v10.877H43.5l-2.053 9.918l-3.081-1.095v2.51c0 2.453-9.318 8.861-14.186 9.888c-4.868-1.027-14.185-7.435-14.185-9.888v-2.51L4.5 30.603l1.35-10.775h4.145V8.95c2.13 1.673 8.443 4.45 14.185.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.012 27.934c1.584 1.44 3.574.595 4.869-1.451c1.239 1.69 3.697 1.126 4.373-.648c.6 1.614 1.952 1.78 2.924 1.78m7.297-1.132l3.634 5.914m-8.008-6.562l1.752 10.593m7.49-8.494c-1.584 1.44-3.573.595-4.868-1.451c-1.24 1.69-3.698 1.126-4.374-.648c-.6 1.614-1.952 1.78-2.923 1.78"/>`),
		g.Group(children),
	)
}