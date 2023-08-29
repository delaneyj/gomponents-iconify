package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ffupdater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.8 40.55c-1.25-2.79-7.18-14.43-15.12-14.43c-6.08 0-7.59 6.28-13.88 6.28a12.42 12.42 0 0 1-4.64-.71a.53.53 0 0 0-.68.31a.64.64 0 0 0 0 .2C7.58 35.77 8.8 42.5 18 42.5c6.67 0 7.31-6.57 14.94-6.57c3.89 0 7.48 3.44 8.69 4.75a.1.1 0 0 0 .15 0a.11.11 0 0 0 .02-.13ZM37.07 19.2L33.23 5.5l-3.34 5.92A15.52 15.52 0 0 0 6.19 24.6a15.2 15.2 0 0 0 .45 3.68h4.12a11.54 11.54 0 0 1 17.13-13.4l-4.38 7.79Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.05 28.28l4.45 1.36l-.53 3.12l2.19.91l-.91 4.61"/>`),
		g.Group(children),
	)
}