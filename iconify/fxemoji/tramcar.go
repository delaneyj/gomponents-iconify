package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tramcar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M176 424c0 26.5-21.5 48-48 48s-48-21.5-48-48m256 0c0 26.5 21.5 48 48 48s48-21.5 48-48" class="st0"/><path d="M0 24h512v8H0z" class="st1"/><path d="m312 27.4l-11-11V16h-90v.4l-11 11L244.6 72L232 84.6L243.4 96L256 83.4L268.6 96L280 84.6L267.4 72L312 27.4zM227.4 32h57.1L256 60.6L227.4 32z" class="st2"/><path d="M444.6 88L437 64.7V64H76v.3L67.4 88H0v312h512V88z" class="st3"/><path d="M0 96h512v224H0z" class="st4"/><path d="M512 496H0v-32h512v32zm0-96H0v32h105.4c3.3 9.3 12.2 16 22.6 16s19.3-6.7 22.6-16h210.7c3.3 9.3 12.2 16 22.6 16s19.3-6.7 22.6-16H512v-32zM368 152H260v144h108V152zm8 0v144h108V152H376zm-240 0H28v144h108V152zm8 0v144h108V152H144z" class="st2"/>`),
		g.Group(children),
	)
}