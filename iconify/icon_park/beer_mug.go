package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeerMug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M16 37.2414C16 39.1724 16.8571 44 28 44C39.1429 44 40 39.1724 40 37.2414C40 35.3103 40 16 40 16H28H16C16 16 16 35.3103 16 37.2414Z"/><path stroke="#fff" stroke-linecap="round" d="M25 23V37"/><path stroke="#fff" stroke-linecap="round" d="M31 23V37"/><path stroke="#000" d="M15.9983 16C15.9983 16 13.9992 11.5 16.9988 9C19.9983 6.5 22.9991 8 22.9991 8C22.9991 8 23.9983 4 27.9991 4C32 4 32.9985 8 32.9985 8C32.9985 8 36.5 6.5 39 9C41.5 11.5 39.9985 16 39.9985 16"/><path stroke="#000" d="M16 19H6C6 19 7 29 7.9991 32C8.99821 35 16 34 16 34"/></g>`),
		g.Group(children),
	)
}