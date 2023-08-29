package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neural(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M20.5 33c1.966 0 2.79-3.538 4.5-4.247c1.772-.735 4.417 1.336 5.753 0c1.336-1.335-.306-4.302.43-6.074C31.89 20.97 35 20.466 35 18.5c0-1.966-3.538-2.44-4.247-4.15c-.735-1.772 1.336-4.767 0-6.103c-1.336-1.335-4.487.735-6.26 0C22.783 7.538 22.466 4 20.5 4c-1.966 0-2.496 3.965-4.206 4.674c-1.772.735-4.711-1.762-6.047-.427c-1.336 1.336.735 3.981 0 5.753C9.538 15.71 6 16.534 6 18.5c0 1.966 3.538 2.47 4.247 4.18c.735 1.771-1.336 4.738 0 6.073c1.336 1.336 4.275-.735 6.047 0c1.71.71 2.24 4.247 4.206 4.247Z" clip-rule="evenodd"/><path stroke-linejoin="round" d="M20 22a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M16.57 30c-1.209 6.8-.494 11.054 2.144 12.766c3.958 2.568 6.734.561 8.183-.91c1.448-1.472 6.153-7.912 10.18-7.912c4.026 0 5.238 3.88 4.856 6c-.254 1.413-.911 2.459-1.97 3.136"/></g>`),
		g.Group(children),
	)
}