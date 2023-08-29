package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleChromeLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 24a83.89 83.89 0 0 1 65.9 32H128a52.05 52.05 0 0 0-46.15 28.07l-17.67-30.6A83.82 83.82 0 0 1 128 44Zm28 84a28 28 0 1 1-28-28a28 28 0 0 1 28 28Zm-112 0a83.41 83.41 0 0 1 6-31.11L83 154c.06.11.14.2.2.3A52 52 0 0 0 128 180q1.19 0 2.34-.06l-17.68 30.63A84.12 84.12 0 0 1 44 128Zm96.05 83.12L173 154c.09-.15.16-.3.24-.46a51.81 51.81 0 0 0-1.46-53.54h35.4a83.95 83.95 0 0 1-67.13 111.12Z"/>`),
		g.Group(children),
	)
}