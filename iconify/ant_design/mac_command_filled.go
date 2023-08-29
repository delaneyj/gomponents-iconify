package ant_design

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MacCommandFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M624 672c0 26.5 21.5 48 48 48s48-21.5 48-48s-21.5-48-48-48h-48v48zm96-320c0-26.5-21.5-48-48-48s-48 21.5-48 48v48h48c26.5 0 48-21.5 48-48z"/><path fill="currentColor" d="M928 64H96c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32zM672 560c61.9 0 112 50.1 112 112s-50.1 112-112 112s-112-50.1-112-112v-48h-96v48c0 61.9-50.1 112-112 112s-112-50.1-112-112s50.1-112 112-112h48v-96h-48c-61.9 0-112-50.1-112-112s50.1-112 112-112s112 50.1 112 112v48h96v-48c0-61.9 50.1-112 112-112s112 50.1 112 112s-50.1 112-112 112h-48v96h48z"/><path fill="currentColor" d="M464 464h96v96h-96zM352 304c-26.5 0-48 21.5-48 48s21.5 48 48 48h48v-48c0-26.5-21.5-48-48-48zm-48 368c0 26.5 21.5 48 48 48s48-21.5 48-48v-48h-48c-26.5 0-48 21.5-48 48z"/>`),
		g.Group(children),
	)
}