package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlurCircular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16A240 240 0 0 0 86.294 425.705A240 240 0 0 0 425.706 86.294A238.432 238.432 0 0 0 256 16Zm147.078 387.078c-81.1 81.1-213.058 81.1-294.157 0s-81.1-213.057 0-294.156a208.238 208.238 0 0 1 294.157 0c81.099 81.099 81.099 213.057 0 294.156Z"/><path fill="currentColor" d="M197.483 242.382a46.332 46.332 0 1 0-32.776-13.555a46.206 46.206 0 0 0 32.776 13.555ZM187.334 185.9a14.354 14.354 0 1 1 0 20.3a14.311 14.311 0 0 1 0-20.3Zm127.183 56.482a46.344 46.344 0 1 0-32.777-79.109a46.332 46.332 0 0 0 32.777 79.108Zm-10.15-56.482a14.354 14.354 0 1 1 0 20.3a14.371 14.371 0 0 1 0-20.3Zm-139.66 97.273a46.353 46.353 0 1 0 65.553 0a46.048 46.048 0 0 0-65.553 0Zm42.926 42.927a14.347 14.347 0 1 1 0-20.3a14.372 14.372 0 0 1 0 20.3Zm74.107-42.927a46.354 46.354 0 1 0 65.553 0a46.406 46.406 0 0 0-65.553 0Zm42.926 42.927a14.354 14.354 0 1 1 4.2-10.15a14.372 14.372 0 0 1-4.2 10.15Z"/><circle cx="314.517" cy="98.5" r="29.637" fill="currentColor" transform="rotate(-9.217 314.534 98.505)"/><circle cx="197.483" cy="98.5" r="29.637" fill="currentColor" transform="rotate(-9.217 197.493 98.505)"/><circle cx="314.517" cy="413.5" r="29.637" fill="currentColor" transform="rotate(-67.5 314.517 413.5)"/><circle cx="197.483" cy="413.5" r="29.637" fill="currentColor" transform="rotate(-67.5 197.483 413.5)"/><circle cx="413.5" cy="314.517" r="29.637" fill="currentColor" transform="rotate(-22.5 413.5 314.518)"/><circle cx="413.5" cy="197.483" r="29.637" fill="currentColor" transform="rotate(-58.283 413.496 197.483)"/><circle cx="98.5" cy="314.517" r="29.637" fill="currentColor"/><circle cx="98.5" cy="197.483" r="29.637" fill="currentColor"/>`),
		g.Group(children),
	)
}