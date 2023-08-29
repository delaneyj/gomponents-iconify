package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 28H72a36 36 0 0 0-36 36v144a20 20 0 0 0 20 20h28a20 20 0 0 0 20-20v-16h48v16a20 20 0 0 0 20 20h28a20 20 0 0 0 20-20V64a36 36 0 0 0-36-36ZM60 168v-56h136v56ZM72 52h112a12 12 0 0 1 12 12v24H60V64a12 12 0 0 1 12-12Zm8 152H60v-12h20Zm96 0v-12h20v12Zm-68-64a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm72 0a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm76-60v24a12 12 0 0 1-24 0V80a12 12 0 0 1 24 0ZM24 80v24a12 12 0 0 1-24 0V80a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}