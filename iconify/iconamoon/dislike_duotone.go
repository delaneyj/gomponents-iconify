package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DislikeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M17.36 4H8v10c.625 0 1.208.312 1.555.832l2.554 3.832A3 3 0 0 0 14.606 20h.213a1 1 0 0 0 .987-1.164L15 14h3.56a2 2 0 0 0 1.962-2.392l-1.2-6A2 2 0 0 0 17.36 4Z" opacity=".16"/><path fill="currentColor" d="m15 14l-.986.164A1 1 0 0 1 15 13v1ZM4 14v1a1 1 0 0 1-1-1h1Zm16.522-2.392l.98-.196l-.98.196ZM6 3h11.36v2H6V3Zm12.56 12H15v-2h3.56v2Zm-2.574-1.164l.806 4.835L14.82 19l-.805-4.836l1.972-.328ZM14.82 21h-.213v-2h.213v2Zm-3.542-1.781l-2.515-3.774l1.664-1.11l2.515 3.774l-1.664 1.11ZM7.93 15H4v-2h3.93v2ZM3 14V6h2v8H3Zm17.302-8.588l1.2 6l-1.961.392l-1.2-6l1.961-.392ZM8.762 15.445A1 1 0 0 0 7.93 15v-2a3 3 0 0 1 2.496 1.336l-1.664 1.11Zm8.03 3.226A2 2 0 0 1 14.82 21v-2l1.973-.329ZM18.56 13a1 1 0 0 0 .981-1.196l1.961-.392A3 3 0 0 1 18.56 15v-2Zm-1.2-10a3 3 0 0 1 2.942 2.412l-1.961.392a1 1 0 0 0-.98-.804V3Zm-2.754 18a4 4 0 0 1-3.329-1.781l1.665-1.11a2 2 0 0 0 1.664.891v2ZM6 5a1 1 0 0 0-1 1H3a3 3 0 0 1 3-3v2Z"/><path stroke="currentColor" stroke-width="2" d="M8 14V4"/></g>`),
		g.Group(children),
	)
}