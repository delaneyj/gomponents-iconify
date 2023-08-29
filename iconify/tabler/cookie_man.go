package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookieMan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2a5 5 0 0 1 2.845 9.112l.147.369l1.755-.803c.969-.443 2.12-.032 2.571.918a1.88 1.88 0 0 1-.787 2.447l-.148.076L16 15.208v2.02l1.426 1.425l.114.125a1.96 1.96 0 0 1-2.762 2.762l-.125-.114l-2.079-2.08l-.114-.124a1.957 1.957 0 0 1-.161-.22H11.7c-.047.075-.101.15-.16.22l-.115.125l-2.08 2.079a1.96 1.96 0 0 1-2.886-2.648l.114-.125L8 17.227v-2.019l-2.383-1.09l-.148-.075a1.88 1.88 0 0 1-.787-2.447c.429-.902 1.489-1.318 2.424-.978l.147.06l1.755.803l.147-.369a5 5 0 0 1-2.15-3.895V7a5 5 0 0 1 5-5zm0 14h.01M12 13h.01M10 7h.01M14 7h.01M12 9h.01"/>`),
		g.Group(children),
	)
}