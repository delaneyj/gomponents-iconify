package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hendroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="15.31" height="27.41" x="16.77" y="6.72" fill="none" stroke="currentColor" stroke-miterlimit="10" rx="1.84" transform="rotate(-18 24.436 20.407)"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M28.66 33.45a3.1 3.1 0 0 0 2.34 2.3c3.51.79 7-7.64 10.43-4.43c3.11 2.93-1.33 6.55-1.33 6.55M28.83 19.7l-2.74 6.39l-5.98-3.56m8.72-2.83a2.83 2.83 0 1 0-5-2.57l-.37.74l-.74-.38a2.83 2.83 0 1 0-2.57 5m10.16-8.72a4.31 4.31 0 0 0 2.27-3.63a10.43 10.43 0 0 0-.85-2.59h2.06A9.08 9.08 0 0 1 40 9.93a6.72 6.72 0 0 1 2.22 5.66c-.44 4.35-1.91 4.82-1.91 4.82s-.21-2.43-2.11-3.12a2.3 2.3 0 0 0-2.7.95a3.79 3.79 0 0 0-1.93-3.07a2.1 2.1 0 0 0-2.62.67ZM15.76 18.5a4.32 4.32 0 0 1-4-1.61A10.54 10.54 0 0 1 11 14.3l-1.71 1.21a9.07 9.07 0 0 0-3.65 5.59A6.74 6.74 0 0 0 7.17 27c2.91 3.26 4.38 2.78 4.38 2.78s-1.26-2.09-.13-3.77a2.29 2.29 0 0 1 2.74-.81a3.8 3.8 0 0 1-.24-3.63a2.1 2.1 0 0 1 2.51-1Z"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M40.07 37.87s2.37 1.37 1.14 3.15c-1.11 1.6-3.58 1.33-5.26 1.48c0-1.69-.6-4.11.86-5.39c1.63-1.44 3.26.76 3.26.76"/>`),
		g.Group(children),
	)
}