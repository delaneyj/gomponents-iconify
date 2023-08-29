package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FemaleSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9B9B9A" d="M51 25.968c0-8.284-6.716-15-15-15c-8.284 0-15 6.716-15 15c0 7.43 5.413 13.597 12.503 14.791v7.152H26v5h7.503v9l4.997.044V52.91H46v-5h-7.5v-7.149l.15-.026C45.67 39.483 51 33.348 51 25.968zm-15 10c-5.523 0-10-4.477-10-10c0-5.522 4.477-10 10-10s10 4.478 10 10c0 5.523-4.477 10-10 10z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M38.5 40.761v7.15H46v5h-7.5v9.044l-4.997-.044v-9H26v-5h7.503v-7.152C26.413 39.565 21 33.397 21 25.97c0-8.285 6.716-15 15-15s15 6.715 15 15c0 7.38-5.33 13.514-12.35 14.766l-.15.026"/><circle cx="36" cy="25.968" r="10"/></g>`),
		g.Group(children),
	)
}