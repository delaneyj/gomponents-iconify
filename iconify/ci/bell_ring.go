package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17v1a3 3 0 1 1-6 0v-1m6 0H9m6 0h3.59c.383 0 .575 0 .73-.052a1 1 0 0 0 .627-.628c.053-.156.053-.348.053-.734c0-.169 0-.253-.014-.334a.998.998 0 0 0-.173-.42c-.048-.067-.108-.127-.227-.246l-.39-.39a.67.67 0 0 1-.196-.474V10a7 7 0 1 0-14 0v3.722a.67.67 0 0 1-.196.474l-.39.39c-.12.12-.179.179-.226.245a1 1 0 0 0-.175.421c-.013.08-.013.165-.013.334c0 .386 0 .578.052.734a1 1 0 0 0 .629.628c.155.052.346.052.729.052H9m9.019-14.986a10 10 0 0 1 3.153 4.002M5.982 2.014a10 10 0 0 0-3.154 4.002"/>`),
		g.Group(children),
	)
}