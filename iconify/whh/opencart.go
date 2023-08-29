package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opencart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768.556 768q23-46 50.5-96t40-73t21.5-45.5t9-38.5q0-52-19.5-89t-47.5-37q-197 0-360-24q-256-37-382-166q-21-24-28-34t-20-33.5t-19.5-54T.556 0q12 33 24.5 58.5t22 40t23.5 27.5t20.5 18t21.5 13t19 11q141 88 765 88q49 0 88.5 40.5t39.5 92.5q0 30-9 61t-30 65t-40 60.5t-55 65.5t-57.5 61.5t-64.5 65.5zm-480 64q40 0 68 28t28 68t-28 68t-68 28t-68-28t-28-68t28-68t68-28zm384 0q40 0 68 28t28 68t-28 68t-68 28t-68-28t-28-68t28-68t68-28z"/>`),
		g.Group(children),
	)
}