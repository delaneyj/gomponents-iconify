package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConstructionWorker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M8.4 45.33a15.6 1.67 0 1 0 31.2 0a15.6 1.67 0 1 0-31.2 0Z" opacity=".15"/><path fill="#ffda8f" d="M5.91 23.72a18.09 18.09 0 1 0 36.18 0a18.09 18.09 0 1 0-36.18 0Z"/><path fill="#ffbe3d" d="M24 5.64a18.09 18.09 0 1 0 18.09 18.08A18.08 18.08 0 0 0 24 5.64Zm0 33.46a16.51 16.51 0 1 1 16.51-16.51A16.51 16.51 0 0 1 24 39.1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M5.91 23.72a18.09 18.09 0 1 0 36.18 0a18.09 18.09 0 1 0-36.18 0Z"/><path fill="#ffe9bd" d="M37.11 32.26c0 .75-1 1.36-2.26 1.36s-2.26-.61-2.26-1.36s1-1.36 2.26-1.36s2.26.61 2.26 1.36Zm-26.22 0a2.26 1.36 0 1 0 4.52 0a2.26 1.36 0 1 0-4.52 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.17 26.51a1.67 1.67 0 0 1-3.34 0m19 0a1.67 1.67 0 1 0 3.34 0m-13.12 10.2h3.9"/><path fill="#ffe500" d="M5.3 20.2a18.7 18.7 0 0 1 37.4 0Z"/><path fill="#fff48c" d="M24 5.7a18.72 18.72 0 0 1 18.23 14.5h.47a18.7 18.7 0 0 0-37.4 0h.47A18.72 18.72 0 0 1 24 5.7Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M5.3 20.2a18.7 18.7 0 0 1 37.4 0Z"/><path fill="#ffe500" d="M3.25 19a1.67 1.67 0 0 1 2-1.64A93.32 93.32 0 0 0 24 19.3a93.32 93.32 0 0 0 18.72-2a1.67 1.67 0 0 1 2 1.64v1.39a1.74 1.74 0 0 1-1.38 1.71A99.76 99.76 0 0 1 24 24.2a99.76 99.76 0 0 1-19.37-2.14a1.74 1.74 0 0 1-1.38-1.71Z"/><path fill="#fff48c" d="M42.72 17.32A93.32 93.32 0 0 1 24 19.3a93.32 93.32 0 0 1-18.72-2a1.67 1.67 0 0 0-2 1.64v1.39a1.63 1.63 0 0 0 .08.44a1.66 1.66 0 0 1 2-1.17a93.32 93.32 0 0 0 18.64 2a93.32 93.32 0 0 0 18.72-2a1.66 1.66 0 0 1 2 1.17a1.63 1.63 0 0 0 .08-.44V19a1.67 1.67 0 0 0-2.08-1.68Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.25 19a1.67 1.67 0 0 1 2-1.64A93.32 93.32 0 0 0 24 19.3a93.32 93.32 0 0 0 18.72-2a1.67 1.67 0 0 1 2 1.64v1.39a1.74 1.74 0 0 1-1.38 1.71A99.76 99.76 0 0 1 24 24.2a99.76 99.76 0 0 1-19.37-2.14a1.74 1.74 0 0 1-1.38-1.71Z"/><path fill="#6dd627" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.21 9.26h-2.09v-2.1A1.12 1.12 0 0 0 25 6.05h-2a1.12 1.12 0 0 0-1.12 1.11v2.1h-2.09a1.1 1.1 0 0 0-1.11 1.11v2a1.11 1.11 0 0 0 1.11 1.11h2.09v2.09A1.12 1.12 0 0 0 23 16.7h2a1.12 1.12 0 0 0 1.12-1.12v-2.09h2.09a1.11 1.11 0 0 0 1.11-1.11v-2a1.1 1.1 0 0 0-1.11-1.12Z"/>`),
		g.Group(children),
	)
}