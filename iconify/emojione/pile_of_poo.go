package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PileOfPoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#89664c" d="M32.2 36.9C15.1 36.9 2 29.8 2 45.7C2 58.5 18.9 62 31.8 62c15 0 30.2-3.5 30.2-16.3c0-15.9-12.2-8.8-29.8-8.8"/><path fill="#9b7861" d="M31.7 20.9c-9.6 0-24.5 1.1-24.5 12.6c0 16.8 49.5 16.8 49.5 0c.1-11.5-14.8-12.6-25-12.6"/><path fill="#a88673" d="M49 16.7C44.4 6.4 22.6 17 26.4 4.5c.6-1.9.5-2.7-.9-2.4C17.7 3.7 12 11.7 14 18.4c6.4 21.4 41.3 12.3 35-1.7"/><path fill="#fff" d="M28.8 34.3c0 4-3.2 7.2-7.2 7.2s-7.2-3.2-7.2-7.2s3.2-7.2 7.2-7.2s7.2 3.2 7.2 7.2"/><circle cx="23.6" cy="34.3" r="3.6" fill="#231f20"/><path fill="#fff" d="M49.6 34.3c0 4-3.2 7.2-7.2 7.2s-7.2-3.2-7.2-7.2s3.2-7.2 7.2-7.2c3.9 0 7.2 3.2 7.2 7.2"/><g fill="#231f20"><circle cx="40.4" cy="34.3" r="3.6"/><path d="M38 50.6c0 3.3-2.7 6-6 6s-6-2.7-6-6s2.7-6 6-6s6 2.7 6 6"/></g>`),
		g.Group(children),
	)
}